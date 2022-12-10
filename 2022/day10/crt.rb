class CPU
  attr_reader :regx, :regx_history, :clock

  def initialize()
    @regx = 1
    @regx_history = []
    @clock = 0
  end

  def noop()
    tick
  end

  def addx(val)
    tick
    tick
    @regx += val
  end

  private
  def tick
    @clock += 1
    @regx_history << @regx
  end
end

class CRT
  attr_reader :lines

  def initialize(cpu)
    @lines = []
    line = ""
    0.upto 5  do |row|
      0.upto 39 do |col|
        tick = 40 * row + col
        if [col - 1, col, col + 1].include? cpu.regx_history[tick]
          line << "X"
        else
          line << "."
        end
      end
      @lines << line
      line = ""
    end
  end
end

cpu = CPU.new
input = ARGV.fetch(0, "input.txt")
File.foreach(input, chomp: true) do |line|
  instruction = line.split
  case instruction[0]
  when "noop"
    cpu.noop
  when "addx"
    cpu.addx instruction[1].to_i
  else
    raise "Invalid instruction"
  end
end

signal_strength = [20, 60, 100, 140, 180, 220].collect do |tick|
  cpu.regx_history[tick - 1] * tick
end
puts "Summed strengths: #{signal_strength.sum}"
puts

crt = CRT.new(cpu)
crt.lines.each {|l| puts "#{l}"}
