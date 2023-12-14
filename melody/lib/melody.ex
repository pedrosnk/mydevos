defmodule Melody do
  alias Melody.Cache

  def get(key) do
    Cache.get(key)
  end

  def set(key, value) do
    Cache.set(key, value)
  end
end
