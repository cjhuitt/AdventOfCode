class Monkey
  attr_reader :name
  def initialize(name, operation)
    @name = name
    @operation = operation
  end

  def to_s
    "#{@name}: #{@operation}"
  end

  def resolve(monkeys)
    return @operation.to_i if @operation.match? /\d+/
    op = @operation.match /(?<a>\w+) (?<op>.) (?<b>\w+)/
    a = monkeys[op[:a]].resolve(monkeys)
    b = monkeys[op[:b]].resolve(monkeys)
    case op[:op]
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
end

monkeys = {}
input = ARGV.fetch(0, "input.txt")
File.foreach(input, chomp: true).with_index do |line, index|
  parts = line.split(": ")
  m = Monkey.new(parts[0], parts[1])
  monkeys[m.name] = m
end

puts "root: #{monkeys['root'].resolve(monkeys)}"
