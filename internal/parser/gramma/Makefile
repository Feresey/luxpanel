gen:
	ragel -Z -G2 -o tokenizer.go tokenizer.go.rl
	goyacc -o gramma.go gramma.y

show:
	ragel -Vp -o gramma.dot ./tokenizer.go.rl
	dot -T svg -o grapth.svg ./gramma.dot