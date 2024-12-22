.PHONY: gen
gen: gen-game gen-combat

gen-%: ragel-% gramma-%

ragel-%:
	ragel -Z -G2 -o ./$(subst ragel-,,$$@)/tokenizer.gen.go ./$(subst ragel-,,$$@)/tokenizer.go.rl

gramma-%:
	go run ./yacc -o ./$$</gramma.go -l -v ./$<</y.output ./$<</gramma.y

# .PHONY: gen-game
# gen-game:
# 	ragel -Z -G2 -o ./game/tokenizer.go ./game/tokenizer.go.rl
# 	go run ./yacc -o ./game/gramma.go -l -v ./game/y.output ./game/gramma.y

# .PHONY: gen-combat
# gen-combat:
# 	ragel -Z -G2 -o ./combat/tokenizer.go ./combat/tokenizer.go.rl
# 	go run ./yacc -o ./combat/gramma.go -l -v ./game/y.output ./combat/gramma.y

# .PHONY: show
# show: show-game show-combat

# .PHONY: show-game
# show-game:
# 	ragel -Vp -o ./game/gramma.dot ./game/tokenizer.go.rl
# 	dot -T svg -o ./game/grapth.svg ./game/gramma.dot

# .PHONY: show-combat
# show-combat:
# 	ragel -Vp -o ./combat/gramma.dot ./combat/tokenizer.go.rl
# 	dot -T svg -o ./combat/grapth.svg ./combat/gramma.dot
