defmodule Puzzle do
  def calculateSum(list1, list2) do
    IO.puts(list1, list2)
  end

  def main do
    case File.read("./input_sample.txt") do
      {:ok, content} ->
        result = String.split(content, ~r/   |\n/)
        IO.puts(result)

      {:error, reason} ->
        IO.puts("Error reading file: #{:file.format_error(reason)}")
    end
  end
end

Puzzle.main()
