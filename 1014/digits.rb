require "prime"

num = gets.to_i
if num == 0
  puts 10
elsif num < 10
  puts num
else
  primes = num.prime_division
  if primes.last[0] > 10
    puts "-1"
  else
    two, three = nil, nil
    if primes.first[0] == 2
      two = primes.first
      three = primes[1] if primes[1] && primes[1][0] == 3
    else
      three = primes.first if primes.first[0] == 3
    end

    four, six, eight, nine = nil, nil, nil, nil
    if three && three[1] > 1
      nine = [9, three[1] / 2]
      three[1] = three[1] % 2
    end

    if two && two[1] >= 3
      eight = [8, two[1] / 3]
      two[1] = two[1] % 3
    end

    if three && three[1] == 1
      if two && two[1] > 0
        two[1] -= 1
        three[1] = 0
        six = [6, 1]
      end
    else
      if two && two[1] == 2
        two[1] = 0
        four = [4, 1]
      end
    end

    primes << four if four
    primes << six if six
    primes << eight if eight
    primes << nine if nine

    result = primes.sort { |x, y| x[0] <=> y[0] }.reduce("")  do |res, prime|
      res << (prime[0].to_s * prime[1])
    end
    puts result
  end
end
