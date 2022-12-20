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
    @position = try if not column.blocked? @offsets.collect {|o| o + try}
    orig = @position
    try = @position + MOVES[:down]
    @position = try if not column.blocked? @offsets.collect {|o| o + try}
    orig.y != @position.y
  end

  def to_s
    "Rock at [#{@offsets.collect{|o| "#{o + @position}"}.join(", ")}]]"
  end
end

class Column
  attr_reader :space
  def initialize
    @space = Array.new(50) {Array.new(7) {false}}
    @top = -1
    @internal_top = -1
    @back = 0
  end

  def blocked?(coords)
    trim_offset = Point.new(0,-@back)
    coords.any? do |coord|
      test = coord + trim_offset
      return true if test.y == -1
      return true if test.x == -1
      return true if test.x == 7
      @space[test.y][test.x]
    end
  end

  def fixate(rock)
    rows = [@internal_top]
    trim_offset = Point.new(0,-@back)
    rock.offsets.each do |offset|
      p = rock.position + offset + trim_offset
      @space[p.y][p.x] = true
      rows << p.y
    end
    @internal_top = rows.max
    @top = @internal_top + @back
    trim_window if @internal_top > 50
    grow_window if @internal_top + 25 > @space.length
  end

  def highest
    @top + 1
  end

  private
  def trim_window
    back = @internal_top
    columns = Array.new(7) {false}
    until columns.count(true) == 7 || back < 0
      columns = columns.zip(@space[back]).collect {|a,b| a or b }
      back -= 1
    end
    return if back < 0
    @back += back
    @space = @space[back..]
    @internal_top -= back
  end

  def grow_window
    @space = @space + Array.new(50) {Array.new(7) {false}}
  end
end

rock_cycle = [:dash, :plus, :reverse_l, :capital_i, :block]
column = Column.new
input = ARGV.fetch(0, "input.txt")
jets = File.open(input, chomp: true, &:readline).strip
rock_count = ARGV.fetch(1, 2022).to_i

jet_index = 0

(0...rock_count).each do |count|
  rock = Rock.new(rock_cycle[count % rock_cycle.length], Point.new(2, column.highest + 3))
  jet_index += 1 while rock.move(jets[jet_index % jets.length], column)
  column.fixate(rock)
  jet_index += 1
  puts "Rock #{count}" if count % 100000 == 0
end

puts "Column reached #{column.highest}"
