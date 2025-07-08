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

case File.read("./input_1.txt") do
  {:ok, content} when content != "" ->
    {oddList, evenList} =
      String.split(content, ~r/\s*|\n/)
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

    value = Puzzle.calculateSum(Enum.sort(oddList), Enum.sort(evenList), 0)

    IO.puts("result: #{value}")

  {:error, reason} ->
    IO.puts("Error reading file: #{:file.format_error(reason)}")
end
