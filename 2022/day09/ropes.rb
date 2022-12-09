require 'set'

class Coordinate
  attr_reader :row, :col

  def initialize( r, c )
    @row = r
    @col = c
  end

  def to_s
    "(#{row}, #{col})"
  end

  def move(dir)
    case dir
    when "U"
      Coordinate.new(@row += 1, @col)
    when "D"
      Coordinate.new(@row -= 1, @col)
    when "R"
      Coordinate.new(@row, @col += 1)
    when "L"
      Coordinate.new(@row, @col -= 1)
    end
  end

  def tail(other)
    dr = @row - other.row
    dc = @col - other.col
    return other if dr.abs <= 1 and dc.abs <= 1
    ddr = [[dr, -1].max, 1].min
    ddc = [[dc, -1].max, 1].min
    Coordinate.new(other.row + ddr, other.col + ddc)
  end
end

head = Coordinate.new(0, 0)
p1_positions = Set.new
p1_tail = Coordinate.new(0, 0)
p2_positions = Set.new
p2_tails = [Coordinate.new(0, 0), Coordinate.new(0, 0), Coordinate.new(0, 0),
         Coordinate.new(0, 0), Coordinate.new(0, 0), Coordinate.new(0, 0),
         Coordinate.new(0, 0), Coordinate.new(0, 0), Coordinate.new(0, 0)]

File.foreach("input.txt", chomp: true).with_index do |line, index|
  parts = line.split
  count = parts[1].to_i
  (0...count).each do
    head = head.move(parts[0])
    p1_tail = head.tail(p1_tail)
    p1_positions.add p1_tail.to_s

    last = head
    (0...9).each do |tailnum|
      p2_tails[tailnum] = last.tail(p2_tails[tailnum])
      last = p2_tails[tailnum]
    end
    p2_positions.add p2_tails.last.to_s
  end
end

puts "Part 1: #{p1_positions.length} unique positions"
puts "Part 2: #{p2_positions.length} unique positions"
