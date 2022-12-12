class Num
  attr_reader :val, :num_factors
  attr_accessor :factor

  def initialize(val, factor, num_factors)
    @val = val
    @factor = factor
    @num_factors = num_factors
  end

  def +(other)
    if other.is_a? Num
      Num.new(@val + other.val, @factor, @num_factors)
    else
      Num.new(@val + other, @factor, @num_factors)
    end
  end

  def *(other)
    ov = if other.is_a? Num
      other.val
    else
      other
    end

    on = if other.is_a? Num
      other.num_factors
    else
      other
    end

    n = @num_factors * on
    val = @val * ov
    n += (val / @factor).floor
    val = val % @factor

    Num.new(val, @factor, n)
  end

  def /(other)
    return self if other == 1
    val = @val / other.to_f
    m = @num_factors / other.to_f
    n = m.floor
    r = m - n
    val = val + (@factor * r).round
    Num.new(val.floor, @factor, n)
  end

  def divisible(other)
    @val % other == 0
  end

  def to_s
    if @num_factors > 0
      "(#{@num_factors} * #{@factor} + #{@val})"
    else
      "#{@val}"
    end
  end
end

class Monkey
  attr_reader :id, :test, :inspections

  def initialize(lines)
    @id = extract_id lines.shift
    @items = extract_items lines.shift
    @operation = extract_op lines.shift
    @test = extract_test lines.shift
    @recipients = {true => extract_recipient(lines.shift),
                   false => extract_recipient(lines.shift)}
    @inspections = 0
  end

  def to_s
    "Monkey #{@id} (#{@inspections} inspections) has items: #{@items.collect{|i| i.to_s}.join(", ")}"
  end

  def inspect_items(monkey_list, relief_factor, debug)
    @items.each do |item|
      val = inspect item
      val = val / relief_factor
      throw_to(monkey_list, val, item, debug)
    end
    @inspections += @items.length
    @items = []
  end

  def catch(item)
    @items << item
  end

  def apply_factor(factor)
    @items.each {|i| i.factor = factor}
  end

  private
  def extract_id(string)
    string.chop.split.last.to_i
  end

  def extract_items(string)
    string.split(":").last.split(", ").collect {|s| Num.new(s.to_i, 0, 0)}
  end

  def extract_op(string)
    string.split("= old ").last.split
  end

  def extract_test(string)
    string.split.last.to_i
  end

  def extract_recipient(string)
    string.split.last.to_i
  end

  def inspect(item)
    case @operation.first
    when "*"
      return item * item if @operation.last == "old"
      item * @operation.last.to_i
    when "+"
      return item + item if @operation.last == "old"
      item + @operation.last.to_i
    else
      item
    end
  end

  def throw_to(monkey_list, item, was, debug)
    monkey_list[@recipients[item.divisible(test)]].catch item
    if debug
      r = monkey_list[@recipients[item.divisible(test)]].id
      puts "\n#{@id} throwing #{item} (was #{was}) to #{r}"
      monkey_list.each {|m| puts "#{m}"}
    end
  end
end

monkeys = []
input = ARGV.fetch(0, "input.txt")
lines = File.readlines(input, chomp: true)
until lines.empty?
  definition = lines.shift(7)
  monkeys << Monkey.new(definition.dup)
end

puts "#{monkeys.length} total monkeys"
puts

def keep_away(monkeys, rounds, relief_factor)
  factor = 1
  monkeys.each {|m| factor *= m.test}
  monkeys.each {|m| m.apply_factor(factor)}
  1.upto rounds do |round|
    monkeys.each do |monkey|
      monkey.inspect_items(monkeys, relief_factor, false)
    end
  end
end

keep_away(monkeys, 20, 3)
inspections = monkeys.collect {|m| m.inspections }.sort
puts "Part 1 monkey business: #{inspections.pop * inspections.pop}"

