def outcome(them, me)
  return 3 if them == "A" and me == "X"
  return 6 if them == "A" and me == "Y"
  return 0 if them == "A" and me == "Z"
  return 0 if them == "B" and me == "X"
  return 3 if them == "B" and me == "Y"
  return 6 if them == "B" and me == "Z"
  return 6 if them == "C" and me == "X"
  return 0 if them == "C" and me == "Y"
  return 3 if them == "C" and me == "Z"
end

def hand(me)
  return 1 if me == "X"
  return 2 if me == "Y"
  return 3 if me == "Z"
end

sum = 0
File.foreach("input.txt").with_index do |line, index|
  them, me = line.split
  sum += outcome(them, me) + hand(me)
end

puts "Total: #{sum}"
