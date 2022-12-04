complete = 0
partial = 0

Pair = Struct.new(:min, :max) do
  def contains(other)
    max >= other.min and min <= other.max and
      min <= other.min and max >= other.max
  end

  def overlaps(other)
    (min..max) === other.min or (min..max) === other.max or
      contains(other) or other.contains(self)
  end
end

File.foreach("input.txt").with_index do |line, index|
  a, b, c, d = line.scan /\d+/
  first, second = Pair.new(a.to_i, b.to_i), Pair.new(c.to_i, d.to_i)
  complete += 1 if first.contains(second) or second.contains(first)
  partial += 1 if first.overlaps(second)
end

puts "#{complete} complete overlaps"
puts "#{partial} partial overlaps"
