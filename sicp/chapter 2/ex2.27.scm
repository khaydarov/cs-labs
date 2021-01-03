(define nil '())

;Iterative approach
(define (deep-reverse items)
	(define (rev-iter lst result)
		(if (null? lst)
				result
				(let ((current (car lst))
							(remain (cdr lst)))
					(if (pair? current)
						(rev-iter remain (cons (deep-reverse current) result))
						(rev-iter remain (cons current result))))))
	(rev-iter items nil))

;Recursive approach
(define (deep-reverse items)
	(if (null? items)
		nil
		(let ((current (car items))
			  (remain (cdr items)))
			(if (pair? current)
				(append (deep-reverse remain)
						(list (deep-reverse current)))
				(append (deep-reverse remain)
						(list current))))))
