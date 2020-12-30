(define (product term a next b)
	(if (> a b) 1 (* (term a) (product term (next a) next b))))

;Variant 1
(define (pi n)

	(define (nominator i)
		(cond
			((= i 1) 2)
			((even? i) (+ i 2))
			(else (+ i 1))))

	(define (denominator i)
		(cond
			((even? i) (+ i 1))
			(else (+ i 2))))

	(* 4 (/ (product nominator 1 inc n) (product denominator 1 inc n))))

;Variant 2
(define (pi n)

	(define (nominator i)
		(cond
			((= i 1) 2)
			((even? i) (+ i 2))
			(else (+ i 1))))

	(define (denominator i)
		(cond
			((even? i) (+ i 1))
			(else (+ i 2))))

	(define (term i)
		(/ (nominator i) (denominator i)))

	(* 4 (product term 1 inc n)))
