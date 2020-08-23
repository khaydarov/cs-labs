<?php

class Node
{
    /**
     * @var mixed
     */
    private $value;

    /**
     * @var Node
     */
    private $left;

    /**
     * @var Node
     */
    private $right;

    /**
     * Node constructor.
     * @param $value
     */
    public function __construct($value)
    {
        $this->value = $value;
    }

    /**
     * @param $value
     */
    public function setValue($value)
    {
        $this->value = $value;
    }

    /**
     * @return mixed
     */
    public function getValue()
    {
        return $this->value;
    }

    /**
     * @param Node $node
     */
    public function setLeft($node)
    {
        $this->left = $node;
    }

    /**
     * @return Node
     */
    public function getLeft()
    {
        return $this->left;
    }

    /**
     * @param Node $node
     */
    public function setRight($node)
    {
        $this->right = $node;
    }

    /**
     * @return Node
     */
    public function getRight()
    {
        return $this->right;
    }
}

class BST
{
    /**
     * @var Node
     */
    private $root;

    /**
     * BST constructor.
     * @param null $value
     */
    public function __construct($value = null)
    {
        if ($value) {
            $this->root = new Node($value);
        }
    }

    /**
     * @return Node
     */
    public function getRoot()
    {
        return $this->root;
    }

    /**
     * @param $value
     * @return Node
     */
    public function search($value)
    {
        $node = $this->root;

        while ($node) {
            if ($node->getValue() > $value) {
                $node = $node->getRight();
            } else if ($node->getValue() < $value) {
                $node = $node->getLeft();
            } else {
                break;
            }
        }

        return $node;
    }

    /**
     * @param $value
     * @return Node
     */
    public function insert($value)
    {
        $node = $this->root;

        if (!$node) {
            $this->root = new Node($value);
        } else {
            while ($node) {
                if ($value > $node->getValue()) {
                    if ($node->getRight()) {
                        $node = $node->getRight();
                    } else {
                        $newNode = new Node($value);
                        $node->setRight($newNode);
                        break;
                    }
                } else if ($value < $node->getValue()) {
                    if ($node->getLeft()) {
                        $node = $node->getLeft();
                    } else {
                        $newNode = new Node($value);
                        $node->setLeft($newNode);
                    }
                } else {
                    break;
                }
            }
        }

        return $node;
    }

    /**
     * @return mixed
     */
    public function min()
    {
        $node = $this->root;
        while ($node->getLeft()) {
            $node = $node->getLeft();
        }

        return $node->getValue();
    }

    /**
     * @return mixed
     */
    public function max()
    {
        $node = $this->root;
        while ($node->getRight()) {
            $node = $node->getRight();
        }

        return $node->getValue();
    }

    /**
     * @param $value
     */
    public function delete($value)
    {
        if (!$this->root) {
            return;
        }

        $this->root = $this->removeNode($this->root, $value);
    }

    /**
     * @param $value
     * @return Node|null
     */
    public function find($value): ?Node
    {
        $current = $this->getRoot();
        while ($current) {
            if ($current->getValue() === $value) {
                return $current;
            }

            if ($current->getValue() < $value) {
                $current = $current->getRight();
            } else if ($current->getValue() > $value) {
                $current = $current->getLeft();
            }
        }

        return $current;
    }

    /**
     * @param Node $node
     * @param mixed $value
     */
    private function removeNode($node, $value)
    {
        if (!$node) {
            return null;
        }

        if ($value < $node->getValue()) {
            $node->setLeft($this->removeNode($node->getLeft(), $value));
            return $node;
        } else if ($value > $node->getValue()) {
            $node->setRight($this->removeNode($node->getRight(), $value));
            return $node;
        } else {
            if (!$node->getLeft() && !$node->getRight()) {
                return null;
            }

            if (!$node->getLeft()) {
                return $node->getRight();
            }

            if (!$node->getRight()) {
                return $node->getLeft();
            }

            $current = $node->getRight();
            while($current->getLeft()) {
                $current = $current->getLeft();
            }

            $value = $current->getValue();

            $node->setValue($value);
            $node->setRight($this->removeNode($current, $value));
        }

        return null;
    }
}

$bst = new BST();
$bst->insert(1);
$bst->insert(2);
$bst->insert(3);
$bst->insert(9);
$bst->insert(4);

//$bst->delete(1);
$bst->delete(9);
$bst->delete(4);
echo $bst->min();
echo $bst->max();
var_dump($bst->find(2));