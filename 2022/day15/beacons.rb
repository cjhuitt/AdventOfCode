require 'set'

Struct.new('Point', :x, :y)

def find_coords(string)
  m = string.match /.*x=(?<x>-?\d+).*y=(?<y>-?\d+)/
  Struct::Point.new(m[:x].to_i, m[:y].to_i)
end

def manhattan(a, b)
  (a.x - b.x).abs + (a.y - b.y).abs
end

def overlaps(a, b)
  return false if a[0] < (b[0] - 1) and a[1] < (b[0] - 1)
  return false if a[0] > (b[1] + 1) and a[1] > (b[1] + 1)
  true
end

def row_span(node_distances, row)
  spans = node_distances.collect do |point, reach|
    vert_discount = (point[1] - row).abs
    next if vert_discount > reach
    dx = (reach - vert_discount)
    [point.x-dx, point.x+dx]
  end.compact
  consolidated = []
  until spans.empty?
    test = spans.pop
    found = consolidated.filter {|span| overlaps(test, span)}
    unless found.empty?
      found.each do |span|
        spans << (span + test).minmax
        consolidated.delete span
      end
    else
      consolidated << test
    end
  end
  consolidated.collect {|span| (span.min..span.max)}
end

beacons = {}
x_span = nil
y_span = nil
input = ARGV.fetch(0, "input.txt")
File.foreach(input, chomp: true).with_index do |line, index|
  parts = line.split(": ")
  sensor = find_coords parts[0]
  beacon = find_coords parts[1]
  beacons[sensor] = beacon
  if x_span and y_span
    x_span = (x_span << sensor.x << beacon.x).minmax
    y_span = (y_span << sensor.y << beacon.y).minmax
  else
    x_span = [sensor.x, beacon.x].sort
    y_span = [sensor.y, beacon.y].sort
  end
end
x_span = [x_span.min - 5, x_span.max + 5]
y_span = [y_span.min - 5, y_span.max + 5]

node_distances = {}
beacons.each_pair do |sensor, beacon|
  node_distances[sensor] = manhattan(sensor, beacon)
end

check_row = ARGV.fetch(1, "2000000").to_i
covered = row_span(node_distances, check_row)
others = Set.new
beacons.collect do |sensor, beacon|
  covered.each do |span|
    others << sensor.x if span.include?(sensor.x) and sensor.y == check_row
    others << beacon.x if span.include?(beacon.x) and beacon.y == check_row
  end
end

cols = 0
covered.each {|s| cols += s.count}
puts "Row #{check_row} has #{cols - others.length} unavailable beacon spots"
