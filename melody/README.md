# Melody

This is a Distributed LRUCache Application. This code is meant to be
embedded into any application that uses the Beam.

## Demo

Checkout this code and open three different terminals. On the first one, start a session
with the command:

```sh
iex -name 'main@127.0.0.1' -S mix
```

You can later start how many other nodes you'd like with dinstinct names.

```sh
iex -name 'one@127.0.0.1' -S mix
iex -name 'two@127.0.0.1' -S mix
```

You can observe the content on each node by running the command o each terminal.

```elixir
 :ets.tab2list(:elements)
```

Start populating and fetching the cache.

```elixir
  Melody.Cache.set("key", "value")
  Melody.Cache.get("key", "value")
```

Or for simplicity run a loop that populates the cache (this commands can run in any node)

```elixir
  for i <- 0..200, do: Melody.Cache.set("#{i}", <<(rem(i, 26) + 65)::utf8>>)
```

Run the command to inspect the nodes data (`:ets.tab2list(:elements)`) on each terminal to
check the data is actually distributed and fetched correctly.
