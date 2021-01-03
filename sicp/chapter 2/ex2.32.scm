(define nil '())
(define (map proc items)
	(if (null? items)
			nil
			(cons (proc (car items))
						(map proc (cdr items)))))

(define (subsets s)
	(if (null? s)
		(list nil)
		(let ((rest (subsets (cdr s))))
			(append rest (map (lambda (r)
							(if (null? r)
								(list (car s))
								(append (list (car s)) r))) rest)))))

;Another solution
(define (subsets s)
	(if (null? s)
		(list nil)
		(let ((rest (subsets (cdr s))))
			(append rest (map (lambda (r)
								(cons (car s) r)) rest)))))
