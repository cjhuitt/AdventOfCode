class Space
  attr_reader :row_span, :col_span, :contents
  def initialize
    @contents = {}
    @row_span = [0,0]
    @col_span = [500,500]
  end

  def add_rock(x, y)
    @contents[[x,y]] = :rock
    @row_span = (@row_span + [y]).minmax
    @col_span = (@col_span + [x]).minmax
  end

  def add_sand(x, y)
    point = [x,y]
    until @contents.has_key? point
      a = [point[0], point[1] + 1]
      b = [point[0] - 1, point[1] + 1]
      c = [point[0] + 1, point[1] + 1]
      if not @contents.has_key? a
        return nil if a[1] > @row_span[1]
        point = a
      elsif not @contents.has_key? b
        return nil if b[1] > @row_span[1]
        return nil if b[0] < @col_span[0]
        point = b
      elsif not @contents.has_key? c
        return nil if c[1] > @row_span[1]
        return nil if c[0] > @col_span[1]
        point = c
      else
        @contents[point] = :sand
        return point
      end
    end
    nil
  end

  def fill_sand(x, y)
    max_row = @row_span.max + 2
    point = [x,y]
    until @contents.has_key? point
      a = [point[0], point[1] + 1]
      b = [point[0] - 1, point[1] + 1]
      c = [point[0] + 1, point[1] + 1]
      if a[1] == max_row
        @contents[point] = :sand
        @col_span = (@col_span + [point[0]]).minmax
        return point
      elsif not @contents.has_key? a
        point = a
      elsif not @contents.has_key? b
        point = b
      elsif not @contents.has_key? c
        point = c
      else
        @contents[point] = :sand
        @col_span = (@col_span + [point[0]]).minmax
        return point
      end
    end
    nil
  end

  def to_s
    lines = []
    max_row = @row_span.max + 1
    min_x = @col_span[0] - 2
    max_x = @col_span[1] + 2
    (0..max_row).each do |r|
      line = ""
      (min_x..max_x).each do |c|
        case @contents[[c,r]]
        when :rock
         line << "#"
        when :sand
         line << "o"
        else
          line << "."
        end
      end
      lines << line
    end
    lines << "#" * (max_x - min_x + 1)
  end
end

cavern = Space.new
total = 0;
input = ARGV.fetch(0, "input.txt")
File.foreach(input, chomp: true).with_index do |line, index|
  last_x = last_y = nil
  line.split(" -> ").each do |node|
    coord = node.split(",")
    x = coord[0].to_i
    y = coord[1].to_i
    if last_x and last_y
      range = [x, last_x].minmax
      (range[0]..range[1]).each {|i| cavern.add_rock(i, y)}
      range = [y, last_y].minmax
      (range[0]..range[1]).each {|i| cavern.add_rock(x, i)}
    end
    cavern.add_rock(x, y)
    last_x = x
    last_y = y
  end
end

puts "#{cavern.contents.length} rocks"
sand = 0
while cavern.add_sand(500, 0)
  sand += 1
end
puts "Part 1: Added #{sand} sand"

until cavern.fill_sand(500, 0) == [500, 0]
  sand += 1
end
# add 1 for the final one placed
sand += 1
puts "Part 2: Added #{sand} sand"

#puts
#cavern.to_s.each {|l| puts l}
