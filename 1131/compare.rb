#(1..1000000000).each do |n|
#  (1..1000000000).each do |k|
(1..100).each do |n|
  (1..100).each do |k|
    File.write("test.txt", "#{n} #{k}")
    mine = `ruby copy.rb < test.txt`.chomp
    ac = `./a.out < test.txt`

    if Random.rand(50) % 41 == 0
      puts "running #{n}, #{k}"
    end

    if mine != ac
      puts n, k, mine, ac
      exit
    end
  end
end
