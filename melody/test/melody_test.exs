defmodule MelodyTest do
  use ExUnit.Case
  doctest Melody

  test "greets the world" do
    assert Melody.hello() == :world
  end
end
