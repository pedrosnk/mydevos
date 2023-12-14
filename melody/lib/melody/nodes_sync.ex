defmodule Melody.NodesSync do
  use GenServer

  alias Melody.NodeList

  def start_link(state) do
    GenServer.start_link(__MODULE__, state, name: __MODULE__)
  end

  @impl GenServer
  def init(state) do
    :net_kernel.monitor_nodes(true)

    if is_main_node?() do
      NodeList.add_node(Node.self())
      {:ok, state}
    else
      {:ok, state, {:continue, :connect_to_main_node}}
    end
  end

  @impl GenServer
  def handle_info({:nodeup, node}, state) do
    IO.puts("node connected #{node}") 
    if is_main_node?() do
      NodeList.add_node(node)
      {:noreply, state, {:continue, :sync_nodes}}
    else
      {:noreply, state}
    end

  end

  @impl GenServer
  def handle_info({:nodedown, node}, state) do
    IO.puts("node disconnected #{node}") 
    if is_main_node?() do
      NodeList.remove_node(node)
      :rpc.multicall(Node.list(), Melody.NodeList, :set_nodes, [NodeList.list_nodes()])
    end
    
    {:noreply, state}
  end

  @impl GenServer
  def handle_continue(:sync_nodes, state) do
    :rpc.multicall(Node.list(), NodeList, :set_nodes, [NodeList.list_nodes()])
    {:noreply, state}
  end

  @impl GenServer
  def handle_continue(:connect_to_main_node, state) do
    main_node = :"main@127.0.0.1"
    unless  main_node in Node.list() do
      Node.connect(main_node)
    end
    {:noreply, state}
  end

  defp is_main_node?() do
    Node.self() == :"main@127.0.0.1"
  end
end
