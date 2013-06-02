(1..20).each do |n|
  (0..15).each do |a|
    result_ac = `echo "#{n} #{a} #{a}" | ./a.out`
    result_wa = `echo "#{n} #{a} #{a}" | ruby 1114.rb`
    puts "#{n} #{a} #{a}:  expected #{result_ac} got #{result_wa}" unless result_ac == result_wa
  end
end
