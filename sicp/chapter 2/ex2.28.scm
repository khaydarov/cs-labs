(define nil '())

;Iterative approach
(define (fringe tree)
	(define (iter items result)
		(if (null? items)
			result
			(let ((current (car items))
				  (remain (cdr items)))
				 (if (pair? current)
					 (iter remain (append result (fringe current)))
					 (iter remain (append result (list current)))))))
	(iter tree nil))

;Recursive approach
(define (fringe tree)
	(if (null? tree)
		nil
		(let ((current (car tree))
			  (remain (cdr tree)))
			 (if (pair? current)
				 (append (fringe current) (fringe remain))
				 (append (list current) (fringe remain))))))
