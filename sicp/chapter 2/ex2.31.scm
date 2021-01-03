(define nil '())
(define (map proc items)
	(if (null? items)
			nil
			(cons (proc (car items))
						(map proc (cdr items)))))

;Without map
(define (tree-map proc tree)
	(cond ((null? tree) nil)
		  ((not (pair? tree)) (proc tree))
		  (else (cons (tree-map proc (car tree))
				      (tree-map proc (cdr tree))))))

;With map
(define (tree-map proc tree)
	(map (lambda (sub-tree)
			(if (pair? sub-tree)
				(tree-map proc sub-tree)
				(proc sub-tree))) tree))
