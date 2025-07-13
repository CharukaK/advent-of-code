defmodule Puzzle do
  def calculateSum(list1, list2, _sum) when length(list1) != length(list2) do
    raise ArgumentError, "List lengths must match"
  end

  def calculateSum([f1 | rest1], [f2 | rest2], sum) do
    calculateSum(rest1, rest2, sum + abs(f2 - f1))
  end

  def calculateSum([], [], sum) do
    sum
  end
end

case File.read("./test_case2.txt") do
  {:ok, content} when content != "" ->
    {oddList, evenList} =
      String.split(content, ~r/\s\s\s|\n/)
      |> Enum.reject(&(&1 == ""))
      |> Enum.map(fn x ->
        case Integer.parse(x) do
          {value, _} -> value
          :error -> raise ArgumentError, "Invalid input: all values must be integers"
        end
      end)
      |> Enum.with_index()
      |> Enum.reduce({[], []}, fn {value, index}, {odd, even} ->
        if rem(index, 2) == 0 do
          {odd, even ++ [value]}
        else
          {odd ++ [value], even}
        end
      end)

    # value = Puzzle.calculateSum(Enum.sort(oddList), Enum.sort(evenList), 0)
    eventFrequency =
      Enum.reduce(evenList, Map.from_keys(oddList, 0), fn x, acc ->
        if Map.has_key?(acc, x) do
          Map.update!(acc, x, &(&1 + 1))
        else
          acc
        end
      end)

    sum = Enum.reduce(oddList, 0, fn x, acc -> acc + x * eventFrequency[x] end)

    # Enum.each(evenList, fn x ->
    #   if Map.has_key?(eventFrequency, x) do
    #     Map.update!(eventFrequency, x, &(&1 + 1))
    #   end
    # end)

    IO.inspect(oddList)
    IO.puts(sum)

  {:error, reason} ->
    IO.puts("Error reading file: #{:file.format_error(reason)}")
end
