class Monkey
  attr_reader :name, :operation
  def initialize(name, operation)
    @name = name
    @operation = operation
  end

  def to_s
    "#{@name}: #{@operation}"
  end

  def resolve(monkeys)
    return @operation.to_i if @operation.to_s.match? /\d+/
    parts = op_parts(monkeys)
    a = parts[0].resolve(monkeys)
    b = parts[2].resolve(monkeys)
    case parts[1]
    when "+"
      a + b
    when "-"
      a - b
    when "*"
      a * b
    when "/"
      a / b
    end
  end

  def op_parts(monkeys)
    op = @operation.match /(?<a>\w+) (?<op>.) (?<b>\w+)/
    [monkeys[op[:a]], op[:op], monkeys[op[:b]]]
  end
end

monkeys = {}
input = ARGV.fetch(0, "input.txt")
File.foreach(input, chomp: true).with_index do |line, index|
  parts = line.split(": ")
  m = Monkey.new(parts[0], parts[1])
  monkeys[m.name] = m
end

root = monkeys['root']
orig = root.resolve(monkeys)
puts "Part 1: root says #{orig}"

def find_human(monkey, monkeys)
  return [] if monkey.name == 'humn'
  op = monkey.operation.match /(?<a>\w+) (?<op>.) (?<b>\w+)/
  return nil if op.nil?
  parents = [monkeys[op[:a]], monkeys[op[:b]]]
  found = parents.collect do |m|
    found = find_human(m, monkeys)
    found.nil? ? nil : [m.name, found]
  end.compact.flatten
  found.empty? ? nil : found
end

path = find_human(root, monkeys)
last = root
value_needed = 0
path.each do |m|
  parts = last.op_parts(monkeys)
  other = parts[0].name == m ? parts[2] : parts[0]
  other_val = other.resolve(monkeys)
  if last == root
    value_needed = other_val
    last = monkeys[m]
    next
  end
  case parts[1]
  when "+"
    value_needed = value_needed - other_val
  when "-"
    if other == parts.first
      value_needed = other_val - value_needed
    else
      value_needed = value_needed + other_val
    end
  when "*"
    value_needed = value_needed / other_val
  when "/"
    if other == parts.first
      value_needed = other_val / value_needed
    else
      value_needed = value_needed * other_val
    end
  end
  last = monkeys[m]
end

puts "Part 2: humn says #{value_needed}"

