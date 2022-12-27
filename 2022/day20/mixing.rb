Struct.new('Node', :value, :prev, :next)

class List
  attr_accessor :length
  def initialize
    @head = @tail = nil
    @length = 0
  end

  def add(node)
    @head = node if @head.nil?
    if @tail.nil?
      @tail = @head
    else
      @tail.next = node
      node.prev = @tail
      @tail = node
    end
    @length += 1
  end

  def move(node)
    return if node.value == 0
    from = remove(node)
    after = advance(from, node.value)
    insert(node, after)
  end

  def remove(node)
    p = node.prev
    n = node.next
    p.next = n unless p.nil?
    n.prev = p unless n.nil?
    node.prev = nil
    node.next = nil
    @head = n if @head == node
    @tail = p if @tail == node
    @length -= 1
    p.nil? ? @tail : p
  end

  def insert(node, after)
    if after.nil?
      @tail.next = node
      node.prev = @tail
      @tail = node
    else
      p = after
      n = after.next
      p.next = node
      node.prev = p
      node.next = n
      n.prev = node unless n.nil?
      @tail = node if @tail == p
    end
    @length += 1
  end

  def advance(node, steps)
    normed = steps % @length
    1.upto(normed).each do
      node = node.next
      node = @head if node.nil?
    end
    node
  end

  def to_s
    n = @head
    lines = [n.value]
    until n == @tail
      n = n.next
      lines << n.value
    end
    lines.join(", ")
  end
end

#multiplier = 1
#loops = 1
multiplier = 811589153
loops = 10
order = []
list = List.new
zero = nil
input = ARGV.fetch(0, "input.txt")
File.foreach(input, chomp: true).with_index do |line, index|
  v = line.to_i * multiplier
  node = Struct::Node.new(v, nil, nil)
  order << node
  list.add(node)
  zero = node if v == 0
end

puts "Start:\t\t#{list}" if list.length < 10
1.upto(loops).each do
  order.each do |node|
    list.move(node)
    puts "#{node.value}:\t\t#{list}" if list.length < 10
  end
end

index = zero
values = (0..2).collect do
  index = list.advance(index, 1000)
  index.value
end
puts "Part 1: #{values} sum to #{values.sum}"
