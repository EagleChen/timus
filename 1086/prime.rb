require 'prime'

primes = Prime.first(15005)
nums = gets.to_i

nums.times do
  index = gets.to_i
  puts primes[index-1]
end
