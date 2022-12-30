class Point
  include Comparable
  attr_accessor :x, :y

  def initialize(x, y)
    @x = x
    @y = y
  end

  def to_s
    "(#{@x},#{@y})"
  end

  def around
    [Point.new(@x - 1, @y - 1), Point.new(@x, @y - 1), Point.new(@x + 1, @y - 1),
     Point.new(@x - 1, @y),                            Point.new(@x + 1, @y),
     Point.new(@x - 1, @y + 1), Point.new(@x, @y + 1), Point.new(@x + 1, @y + 1)]
  end

  def <=>(other)
    return nil if other.nil?
    return -1 if @x < other.x
    return 1 if @x > other.x
    return -1 if @y < other.y
    return 1 if @y > other.y
    0
  end

  def +(other)
    Point.new(@x + other.x, @y + other.y)
  end
end

elves = []
input = ARGV.fetch(0, "input.txt")
File.foreach(input, chomp: true).with_index do |line, y|
  line.each_char.with_index do |ch, x|
    elves << Point.new(x, y) if ch == "#"
  end
end

def bounds(elves)
  x_bounds = Array.new(2, elves.first.x)
  y_bounds = Array.new(2, elves.first.y)
  elves.each do |elf|
    x_bounds = (x_bounds << elf.x).minmax
    y_bounds = (y_bounds << elf.y).minmax
  end
  [x_bounds, y_bounds]
end

def empty_spots(elves)
  b = bounds(elves)
  width = b.first.max - b.first.min + 1
  height = b.last.max - b.last.min + 1
  width * height - elves.count
end

def clear?(elf, places, elves)
  near = places.filter {|p| elves.include? p}
  near.empty?
end

def print_grid(label, points)
  puts label
  -2.upto(9).each do |y|
    -3.upto(10).each do |x|
      print points.include?(Point.new(x, y)) ? "#" : "."
    end
    print "\n"
  end
  puts
end
print_grid("Start", elves) if elves.count < 30

test = {north: [Point.new(-1, -1), Point.new( 0, -1), Point.new( 1, -1)],
        south: [Point.new(-1,  1), Point.new( 0,  1), Point.new( 1,  1)],
        west:  [Point.new( 1, -1), Point.new( 1,  0), Point.new( 1,  1)],
        east:  [Point.new(-1, -1), Point.new(-1,  0), Point.new(-1,  1)]}
move = {north: Point.new(0, -1), south: Point.new(0, 1), west: Point.new(1, 0),
        east: Point.new(-1, 0)}
test_order = [:north, :south, :east, :west]

1.upto(10).each do |round|
  proposed = []
  elves.each do |elf|
    propose = nil
    propose = elf if clear?(elf, elf.around, elves)
    test_order.each do |try|
      view = test[try].collect{ |delta| elf + delta }
      if propose.nil? and clear?(elf, view, elves)
        propose = elf + move[try]
        break
      end
    end
    proposed << (propose.nil? ? elf : propose)
  end
  new = []
  elves.zip(proposed).each do |current, want|
    if current != want and proposed.count(want) == 1
      new << want
    else
      new << current
    end
  end
  elves = new
  test_order.rotate!
  print_grid("Round #{round}", elves) if elves.count < 30
end

puts "Part 1: #{empty_spots(elves)} empty positions"
