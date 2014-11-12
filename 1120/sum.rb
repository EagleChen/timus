require "prime"

num = gets.to_i * 2

max = (num ** 0.5).floor
max.downto(1).each do |p|
  if (num % p == 0) && (p % 2 != num/p % 2)
    a = (num/p - p + 1) / 2
    puts "#{a} #{p}"
    break
  end
end
