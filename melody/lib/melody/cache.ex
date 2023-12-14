defmodule Melody.Cache do
  use GenServer

  alias Melody.NodeList

  @default_init_args [
    limit: 50,
    last_used_expiration: {-2, :hour},
    check_expiraton_interval: :timer.minutes(15)
  ]

  def start_link(state \\ []) do
    state = Keyword.merge(@default_init_args, state)
    GenServer.start_link(__MODULE__, state, name: __MODULE__)
  end

  def set(key, value) do
    node = node_for_key(key)

    cond do
      Node.self() == node ->
        GenServer.call(__MODULE__, {:set, key, value})

      true ->
        :rpc.call(node, __MODULE__, :set, [key, value])
    end
  end

  def get(key) do
    node = node_for_key(key)

    cond do
      Node.self() == node ->
        GenServer.call(__MODULE__, {:get, key})

      true ->
        :rpc.call(node, __MODULE__, :get, [key])
    end
  end

  defp ensure_elements_table_exist() do
    :ets.new(
      :elements,
      [:named_table, :public, :set]
    )
  end

  defp ensure_last_used_table_exist() do
    :ets.new(
      :last_used,
      [:named_table, :public, :ordered_set]
    )
  end

  @impl GenServer
  def init(state) do
    ensure_elements_table_exist()
    ensure_last_used_table_exist()

    Process.send_after(self(), :remove_old_entries, state[:check_expiraton_interval])

    {:ok, state}
  end

  @impl GenServer
  def handle_info(:remove_old_entries, state) do
    expiration_time = expiration_time_in_μs(state[:last_used_expiration])

    expired_entries =
      :ets.select(:last_used, [{{:"$1", :"$2"}, [{:<, :"$1", expiration_time}], [:"$$"]}])

    for [expired_ts, expired_key] <- expired_entries do
      :ets.delete(:last_used, expired_ts)
      :ets.delete(:elements, expired_key)
    end

    Process.send_after(self(), :remove_old_entries, :timer.minutes(15))
    {:noreply, state}
  end

  @impl GenServer
  def handle_call({:set, key, value}, _from, state) do
    element = %{
      time: current_time_in_μs(),
      value: value
    }

    case :ets.lookup(:elements, key) do
      [{^key, prev_elem}] ->
        :ets.delete(:last_used, prev_elem.time)

      [] ->
        table_size = :ets.info(:elements)[:size]

        if table_size >= state[:limit] do
          timestamp_for_deletion = :ets.first(:last_used)
          [{_, key_for_deletion}] = :ets.lookup(:last_used, timestamp_for_deletion)
          :ets.delete(:elements, key_for_deletion)
          :ets.delete(:last_used, timestamp_for_deletion)
        end
    end

    :ets.insert(:elements, {key, element})
    :ets.insert(:last_used, {element.time, key})

    {:reply, :ok, state}
  end

  @impl GenServer
  def handle_call({:get, key}, _from, state) do
    case :ets.lookup(:elements, key) do
      [] ->
        {:reply, {:error, :not_found}, state}

      [{^key, elem}] ->
        new_elem = %{elem | time: current_time_in_μs()}
        :ets.delete(:last_used, elem.time)
        :ets.insert(:last_used, {new_elem.time, key})
        :ets.insert(:elements, {key, new_elem})
        {:reply, {:ok, elem.value}, state}
    end
  end

  defp expiration_time_in_μs({amount, unit}) do
    DateTime.utc_now() |> DateTime.add(amount, unit) |> DateTime.to_unix(:microsecond)
  end

  defp current_time_in_μs() do
    DateTime.utc_now() |> DateTime.to_unix(:microsecond)
  end

  defp node_for_key(key) do
    key
    |> then(fn k -> :crypto.hash(:md5, k) end)
    |> :erlang.binary_to_list()
    |> Enum.sum()
    |> then(&(rem(&1, NodeList.num_partitions()) + 1))
    |> NodeList.node_on_partition()
  end
end
