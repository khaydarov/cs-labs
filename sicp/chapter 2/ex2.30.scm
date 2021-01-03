(define nil '())
(define (map proc items)
	(if (null? items)
			nil
			(cons (proc (car items))
						(map proc (cdr items)))))

;With map
(define (square-tree tree)
	(map (lambda (sub-tree)
		(if (pair? sub-tree)
			(square-tree sub-tree)
			(* sub-tree sub-tree))) tree))

;Without map
(define (square-tree tree)
	(cond ((null? tree) nil)
		  ((not (pair? tree)) (* tree tree))
		  (else (cons (square-tree (car tree))
					   (square-tree (cdr tree))))))
