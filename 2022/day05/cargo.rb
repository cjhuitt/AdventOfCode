stacks = [Array.new, Array.new, Array.new, Array.new, Array.new,
          Array.new, Array.new, Array.new, Array.new]
initializing = true

File.foreach("input.txt").with_index do |line, index|
  unless line.strip.empty?
    if initializing
      vals = line.scan /.(.)...(.)...(.)...(.)...(.)...(.)...(.)...(.)...(.)./m
      vals.flatten.each.with_index do |val, index|
        stacks[index].push val if not val.strip.empty?
      end
    else
      a, b, c = line.scan /(\d+)/
      count, from, to = a.join.to_i, b.join.to_i - 1, c.join.to_i - 1
      shift = stacks[from].pop count
      shift.reverse.each do |crate|
        stacks[to].push crate
      end
    end
  else
    initializing = false
    stacks.each do |stack|
      stack.reverse!
    end
  end
end

puts "Tops: #{stacks.map{|s| s.last}.join}"
