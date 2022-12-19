Struct.new('Valve', :id, :flow, :connections)

def parse(line)
  m = line.match /Valve (?<id>\w+) .*=(?<flow>\d+);.* valves? (?<conns>.*)/
  connections = m[:conns].split(", ")
  Struct::Valve.new(m[:id], m[:flow].to_i, connections)
end

valves = {}
useful = []
input = ARGV.fetch(0, "input.txt")
File.foreach(input, chomp: true).with_index do |line, index|
  valve = parse(line)
  valves[valve.id] = valve
  useful << valve.id if valve.flow > 0
end

puts "#{valves.count} valves, #{useful.count} useful"

def shuffled(placed, remaining)
  return [placed] if remaining.empty?
  remaining.collect do |l|
    remains = remaining.dup
    remains.delete l
    shuffled(placed.dup << l, remains)
  end.flatten(1)
end
valve_orders = shuffled([], useful)
puts
puts "#{valve_orders.count} valve order options"

def shortest_path_between(valves, goal, current_location, visited)
  return [current_location, goal] if valves[current_location].connections.include? goal
  v = visited.dup() << current_location
  shortest = nil
  valves[current_location].connections.each do |node|
    next if visited.include? node
    path = shortest_path_between(valves, goal, node, v)
    next if path.nil?
    shortest = path if shortest.nil?
    shortest = path if path.length < shortest.length
  end
  return nil if shortest.nil?
  [current_location] + shortest
end

def full_path(valves, waypoints)
  path = []
  last = "AA"
  waypoints.each do |node|
    path << shortest_path_between(valves, node, last, [])[1..]
    last = node
  end
  path
end

routes = valve_orders.collect {|o| full_path(valves, o)}

def score(valves, path)
  minutes_left = 30 # Account for starting at "AA"
  current_flow = 0
  score = 0
  path.each do |leg|
    leg.each do |node|
      score += current_flow
      minutes_left -= 1
      break if minutes_left <= 0
    end
    last = leg.last
    score += current_flow
    minutes_left -= 1
    current_flow += valves[last].flow
  end
  score += current_flow * minutes_left
end

best_score = 0
best_route = []
routes.each do |route|
  s = score(valves, route)
  if s > best_score
    best_score = s
    best_route = route
  end
end
puts "Best score: #{best_score} (route #{best_route})"

