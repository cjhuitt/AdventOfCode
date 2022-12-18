class Point
  attr_accessor :x, :y
  def initialize(x, y)
    @x = x
    @y = y
  end

  def +(other)
    Point.new(@x + other.x, @y + other.y)
  end

  def to_s
    "(#{@x},#{@y})"
  end
end

class Rock
  attr_accessor :position
  attr_reader :offsets
  SHAPES = {dash: [Point.new(0,0),Point.new(1,0),Point.new(2,0),Point.new(3,0)],
            plus: [Point.new(0,1),Point.new(1,2),Point.new(1,1),Point.new(1,0),Point.new(2,1)],
            reverse_l: [Point.new(0,0),Point.new(1,0),Point.new(2,2),Point.new(2,1),Point.new(2,0)],
            capital_i: [Point.new(0,3),Point.new(0,2),Point.new(0,1),Point.new(0,0)],
            block: [Point.new(0,1),Point.new(0,0),Point.new(1,1),Point.new(1,0)]
  }
  MOVES = {"<" => Point.new(-1,0), ">" => Point.new(1, 0), down: Point.new(0, -1)}

  def initialize(type, origin)
    @offsets = SHAPES[type]
    @position = origin
  end

  def move(jet, column)
    try = @position + MOVES[jet]
    @position = try if not blocked?(try, column)
    orig = @position
    try = @position + MOVES[:down]
    @position = try if not blocked?(try, column)
    orig.y != @position.y
  end

  def to_s
    "Rock at [#{@offsets.collect{|o| "#{o + @position}"}.join(", ")}]]"
  end

  private
  def blocked?(pos, column)
    @offsets.any? do |offset|
      check = offset + pos
      return true if check.y == -1
      return true if check.x == -1
      return true if check.x == 7
      column[check.y][check.x]
    end
  end
end

rock_cycle = [:dash, :plus, :reverse_l, :capital_i, :block]
column = Array.new(11000) {Array.new(7) {false}}
input = ARGV.fetch(0, "input.txt")
jets = File.open(input, chomp: true, &:readline).strip

def print_column(column, lines)
  (0..lines).each do |row|
    print "|#{column[lines - row].collect {|v| v ? "#" : "."}.join}|\n"
  end
  print "+-------+"
end

def fixate(rock, column)
  rows = []
  rock.offsets.each do |offset|
    p = rock.position + offset
    column[p.y][p.x] = true
    rows << p.y
  end
  rows
end

top = -1
jet_index = 0

(0..2021).each do |count|
  rock = Rock.new(rock_cycle[count % rock_cycle.length], Point.new(2, top + 4))
  while rock.move(jets[jet_index % jets.length], column)
    jet_index += 1
  end
  jet_index += 1
  top = [top, fixate(rock, column).max].max
end

puts "Column reached #{top + 1}"
