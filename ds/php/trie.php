<?php

/**
 * Class Node
 */
class Node
{
    public $isEnd = false;
    public $keys = [];
}

/**
 * Class Trie
 */
class Trie
{
    /**
     * @var Node
     */
    private $root;

    /**
     * Trie constructor.
     */
    public function __construct()
    {
        $this->root = new Node();
    }

    /**
     * @param $word
     * @param $node
     */
    public function add($word, $node = null): void
    {
        if (!$node) {
            $node = $this->root;
        }

        if (strlen($word) === 0) {
            $node->isEnd = true;
        } else if (!isset($node->keys[$word[0]])) {
            $node->keys[$word[0]] = new Node();
            $this->add(substr($word, 1), $node->keys[$word[0]]);
        } else {
            $this->add(substr($word, 1), $node->keys[$word[0]]);
        }
    }

    /**
     * @param $word
     * @return bool
     */
    public function hasWord($word): bool
    {
        $node = $this->root;

        while (strlen($word) > 1) {
            if (!isset($node->keys[$word[0]])) {
                return false;
            }

            $node = $node->keys[$word[0]];
            $word = substr($word, 1);
        }

        return isset($node->keys[$word]) && $node->keys[$word]->isEnd;
    }

    /**
     * @return array
     */
    public function printWords(): array
    {
        $dictionary = [];
        $this->collectWord($this->root, '', $dictionary);
        return $dictionary;
    }

    /**
     * @param $char
     * @return array
     */
    public function findWords($string): array
    {
        $word = $string;
        $words = [];
        $root = $this->root;

        if (strlen($string) === 1 && !isset($root->keys[$string])) {
            return [];
        }

        while (strlen($string) > 1) {
            $char = $string[0];
            if (isset($root->keys[$char])) {
                $root = $root->keys[$char];
            }
            $string = substr($string, 1);
        }
        $this->collectWord($root->keys[$string], $word, $words);

        return $words;
    }

    /**
     * @param $node
     * @param string $string
     * @param array $dictionary
     */
    public function collectWord($node, $string = '', array &$dictionary = []): void
    {
        if (!empty($node->keys)) {
            foreach ($node->keys as $char => $subNode) {
                $this->collectWord($subNode, $string . $char, $dictionary);
            }

            if ($node->isEnd) {
                $dictionary[] = $string;
            }
        } else {
            $dictionary[] = $string;
        }
    }
}

$trie = new Trie();
$trie->add('hello');
$trie->add('hi');
$trie->add('bye');
$trie->add('bye-bye');
$trie->add('beep');

//print_r($trie->printWords());
//print_r($trie->findWords('t'));