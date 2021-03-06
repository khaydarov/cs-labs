(define (simpleRule f a b n)
	(define h (/ (- b a) n))
	(define (term i)
		(* (f (+ a (* i (/ (- b a) n))))
			(multiple i)))
	(define (multiple x)
		(cond ((or (= x 0) (= x n)) 1)
				((even? x) 2)
				((odd? x) 4)))
	(define (inc a)
		(+ a 1))
	(* (sum term 0 inc n) (/ h 3)))
