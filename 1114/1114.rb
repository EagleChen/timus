boxes, balls, useless = gets.split(" ").map { |s| s.to_i }

def one_kind(boxes, balls)
  all = boxes + balls
  pick = boxes - 1
  (1..pick).reduce(1) { |result, i| result * (all-i) / i }
end

puts one_kind(boxes+1, balls) ** 2 if balls > 0
puts 1 if balls == 0
