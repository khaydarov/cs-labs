;Recursive
(define (* a b)
	(define (double a)
		(+ a a))
	(cond ((= b 0) 0)
			((even? b) (double (* a (/ b 2))))
			(else (+ a (* a (- b 1))))))

;Iterative
(define (* a b)
	(define (double a)
		(+ a a))
	(define (iter a b product)
		(cond ((= b 0) product)
				((even? b) (iter a (/ b 2) (double product)))
				(else (iter a (- b 1) (+ a product)))))
	(iter a b 0))