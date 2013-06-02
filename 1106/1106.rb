n = gets.to_i

teams = {}
count = 0
(1..n).each do |node|
  node = node.to_s
  connections = gets.split(" ")
  unless teams.key? node
    connections.delete_at(connections.size - 1)
    if connections.size == 0
      count = 0
      break
    end
    connections.each do |other_node|
      if teams.key? other_node
        teams[node] = 3 - teams[other_node]
        count += 1 if teams[node] == 1
        break
      end
    end
    unless teams.key? node
      teams[node] = 1
      count += 1
      connections.each do |other_node|
        teams[other_node] = 2
      end
    end
  end
end

puts count
puts teams.select { |k,v| v == 1 }.to_a.map { |o| o[0] }.join(" ") if count > 0
