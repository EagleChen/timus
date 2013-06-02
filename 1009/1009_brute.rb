n = gets.split()[0].to_i
k = gets.split()[0].to_i

def start_num(n, k)
  k ** (n-1)
end

def end_num(n, k)
  k ** n -1
end

def cal(num, k)
  previous_is_zero = false
  loop do
    quotient = num / k
    remainder = num % k
    if remainder == 0
      return 0 if previous_is_zero
      previous_is_zero = true
    else
      previous_is_zero = false
    end
    break if quotient == 0 
    num = quotient
  end
  1
end

puts (start_num(n, k)..end_num(n, k)).reduce(0) { |result, num| result + cal(num, k) }
