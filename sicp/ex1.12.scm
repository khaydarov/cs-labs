;Recursive
(define (pascal row col)
	(cond ((or (= row 0) (= col 0)) (error "must be greater than 0"))
		  ((> col row) 0)
		  ((or (= col 1) (= col row)) 1)
  		  (else (+ (pascal (- row 1) col)
					(pascal (- row 1) (- col 1))))))

;Binomial
(define (pascal-triangle row col)
	(define (factorial n)
		(define (iter n product)
			(if (< n 2)
				product
				(iter (- n 1) (* n product))))
		(iter n 1))
	(/ (factorial (- row 1)) (* (factorial (- col 1)) (factorial (- (- row 1) (- col 1))))))
