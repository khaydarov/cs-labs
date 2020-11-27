(define (accumulate combiner null-value term a next b)
	(if (> a b)
		null-value
		(combiner (term a) (accumulate combiner null-value term (next a) next b))))

;Sigma
(define (sum a b)
	(accumulate + 0 identity a inc b))

;Factorial
(define (factorial n)
	(accumulate * 1 identity 1 inc n))
