import { describe, expect, it } from 'vitest';

import { filterTree, mapTree, traverseTreeValues } from './tree';

describe('traverseTreeValues', () => {
  interface Node {
    children?: Node[];
    name: string;
  }

  type NodeValue = string;

  const sampleTree: Node[] = [
    {
      children: [
        { name: 'B' },
        {
          children: [{ name: 'D' }, { name: 'E' }],
          name: 'C',
        },
      ],
      name: 'A',
    },
    {
      children: [
        { name: 'G' },
        {
          children: [{ name: 'I' }],
          name: 'H',
        },
      ],
      name: 'F',
    },
  ];

  it('traverses tree and returns all node values', () => {
    const values = traverseTreeValues<Node, NodeValue>(
      sampleTree,
      (node) => node.name,
      {
        childProps: 'children',
      },
    );
    expect(values).toEqual(['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I']);
  });

  it('handles empty tree', () => {
    const values = traverseTreeValues<Node, NodeValue>([], (node) => node.name);
    expect(values).toEqual([]);
  });

  it('handles tree with only root node', () => {
    const rootNode = { name: 'A' };
    const values = traverseTreeValues<Node, NodeValue>(
      [rootNode],
      (node) => node.name,
    );
    expect(values).toEqual(['A']);
  });

  it('handles tree with only leaf nodes', () => {
    const leafNodes = [{ name: 'A' }, { name: 'B' }, { name: 'C' }];
    const values = traverseTreeValues<Node, NodeValue>(
      leafNodes,
      (node) => node.name,
    );
    expect(values).toEqual(['A', 'B', 'C']);
  });
});

describe('filterTree', () => {
  const tree = [
    {
      children: [
        { id: 2 },
        { children: [{ id: 4 }, { id: 5 }, { id: 6 }], id: 3 },
        { id: 7 },
      ],
      id: 1,
    },
    { children: [{ id: 9 }, { id: 10 }], id: 8 },
    { id: 11 },
  ];

  it('should return all nodes when condition is always true', () => {
    const result = filterTree(tree, () => true, { childProps: 'children' });
    expect(result).toEqual(tree);
  });

  it('should return only root nodes when condition is always false', () => {
    const result = filterTree(tree, () => false);
    expect(result).toEqual([]);
  });

  it('should return nodes with even id values', () => {
    const result = filterTree(tree, (node) => node.id % 2 === 0);
    expect(result).toEqual([{ children: [{ id: 10 }], id: 8 }]);
  });

  it('should return nodes with odd id values and their ancestors', () => {
    const result = filterTree(tree, (node) => node.id % 2 === 1);
    expect(result).toEqual([
      {
        children: [{ children: [{ id: 5 }], id: 3 }, { id: 7 }],
        id: 1,
      },
      { id: 11 },
    ]);
  });

  it('should return nodes with "leaf" in their name', () => {
    const tree = [
      {
        children: [
          { name: 'leaf 1' },
          {
            children: [{ name: 'leaf 2' }, { name: 'leaf 3' }],
            name: 'branch',
          },
          { name: 'leaf 4' },
        ],
        name: 'root',
      },
    ];
    const result = filterTree(
      tree,
      (node) => node.name.includes('leaf') || node.name === 'root',
    );
    expect(result).toEqual([
      {
        children: [{ name: 'leaf 1' }, { name: 'leaf 4' }],
        name: 'root',
      },
    ]);
  });
});

describe('mapTree', () => {
  it('map infinite depth tree using mapTree', () => {
    const tree = [
      {
        children: [
          { id: 2, name: 'node2' },
          { id: 3, name: 'node3' },
          {
            children: [
              {
                children: [
                  { id: 6, name: 'node6' },
                  { id: 7, name: 'node7' },
                ],
                id: 5,
                name: 'node5',
              },
              { id: 8, name: 'node8' },
            ],
            id: 4,
            name: 'node4',
          },
        ],
        id: 1,
        name: 'node1',
      },
    ];
    const newTree = mapTree(tree, (node) => ({
      ...node,
      name: `${node.name}-new`,
    }));

    expect(newTree).toEqual([
      {
        children: [
          { id: 2, name: 'node2-new' },
          { id: 3, name: 'node3-new' },
          {
            children: [
              {
                children: [
                  { id: 6, name: 'node6-new' },
                  { id: 7, name: 'node7-new' },
                ],
                id: 5,
                name: 'node5-new',
              },
              { id: 8, name: 'node8-new' },
            ],
            id: 4,
            name: 'node4-new',
          },
        ],
        id: 1,
        name: 'node1-new',
      },
    ]);
  });
});
