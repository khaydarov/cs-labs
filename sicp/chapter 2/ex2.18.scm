;Iterative
(define (reverse items)
   (define (iter items result)
     (if (null? items)
         result
         (iter (cdr items) (cons (car items) result))))

   (iter items nil))

;Iterative-2
(define (reverse lst)
	(define (do-rev n)
		(if (< n 0)
			nil
			(cons (list-ref lst n) (do-rev (- n 1)))))
	(do-rev (- (length lst) 1)))

;Recursive
(define (reverse lst1)
	(if (null? lst1)
		nil
		(append (reverse (cdr lst1))
				(list (car lst1)))))
