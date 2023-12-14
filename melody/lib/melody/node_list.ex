defmodule Melody.NodeList do
  use Agent

  # [{node, partition}]

  def start_link(initial_value \\ []) do
    Agent.start_link(fn -> initial_value end, name: __MODULE__)
  end

  def set_nodes(nodes) do
    Agent.update(__MODULE__, fn _ -> nodes end)
  end

  def add_node(node) do
    Agent.update(__MODULE__, fn nodes ->
      found_node =
        Enum.find(nodes, fn
          {^node, _} -> true
          _ -> false
        end)

      if found_node do
        nodes
      else
        partition = Enum.count(nodes) + 1
        nodes ++ [{node, partition}]
      end
    end)
  end

  def remove_node(node) do
    Agent.update(__MODULE__, fn nodes ->
      for {n, _} <- nodes, n != node do
        n
      end
      |> Enum.with_index()
      |> Enum.map(fn {node, i} -> {node, i + 1} end)
    end)
  end

  def num_partitions() do
    Agent.get(__MODULE__, &Enum.count(&1))
  end

  def node_on_partition(number) do
    Agent.get(__MODULE__, fn nodes ->
      Enum.find_value(nodes, fn
        {node, ^number} -> node
        _ -> nil
      end)
    end)
  end

  def list_nodes() do
    Agent.get(__MODULE__, fn nodes ->
      nodes
    end)
  end
end
