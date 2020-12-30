;tree-recursive procedure
(define (f n)
	(cond 	((< n 3) n)
			(else (+ (f (- n 1)) (* 2 (f (- n 2))) (* 3 (f (- n 3)))))
	)
)
		
;iterative procedure
(define (f1 n)
	(cond 	((= n 1) 1)
			((= n 2) 2)
			(else (f-iter 0 1 2 (- n 2)))))