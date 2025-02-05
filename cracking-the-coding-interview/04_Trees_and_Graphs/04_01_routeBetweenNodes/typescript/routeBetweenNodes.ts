type GraphMatrix = number[][] 



function parseMermaidToAdjacencyList(mermaid: string): Record<string, string[]> {
    const adjacencyList: Record<string, string[]> = {};
    const nodePattern = /^(\s*)(\w+)\s*\(\((.*?)\)\)/;
    const directedEdgePattern = /^(\s*)(\w+)\s*-->\s*([\w\s&]+)/;
    const undirectedEdgePattern = /^(\s*)(\w+)\s*(---|<-->)\s*([\w\s&]+)/;

    const lines = mermaid.split("\n");

    for (const line of lines) {
        let nodeMatch = line.match(nodePattern);
        if (nodeMatch) {
            const nodeName = nodeMatch[2];
            if (!(nodeName in adjacencyList)) {
                adjacencyList[nodeName] = [];
            }
            continue;
        }

        let directedMatch = line.match(directedEdgePattern);
        if (directedMatch) {
            const fromNode = directedMatch[2];
            const toNodes = directedMatch[3].split("&").map(node => node.trim());

            if (!(fromNode in adjacencyList)) {
                adjacencyList[fromNode] = [];
            }

            for (const toNode of toNodes) {
                if (!(toNode in adjacencyList)) {
                    adjacencyList[toNode] = [];
                }
                adjacencyList[fromNode].push(toNode);
            }
            continue;
        }

        let undirectedMatch = line.match(undirectedEdgePattern);
        if (undirectedMatch) {
            const nodeA = undirectedMatch[2];
            const nodeBList = undirectedMatch[4].split("&").map(node => node.trim());

            if (!(nodeA in adjacencyList)) {
                adjacencyList[nodeA] = [];
            }

            for (const nodeB of nodeBList) {
                if (!(nodeB in adjacencyList)) {
                    adjacencyList[nodeB] = [];
                }
                adjacencyList[nodeA].push(nodeB);
                adjacencyList[nodeB].push(nodeA); // Add back reference for undirected
            }
        }
    }

    return adjacencyList;
}

type AdjacencyList = Record<string, string[]>;

function adjacencyListToMermaid(
    adjList: Record<string, string[]>,
    direction: string = "LR",
    groupLimit: number = 3
): string {
    let mermaid = `\`\`\`mermaid\ngraph ${direction}\n`;

    const processedEdges = new Set<string>();

    // Declare nodes first
    for (const node in adjList) {
        mermaid += `    ${node}(( ${node} ))\n`;
    }

    // Process edges
    for (const node in adjList) {
        const neighbors = adjList[node];

        const directedEdges: string[] = [];
        const undirectedEdges: string[] = [];

        for (const neighbor of neighbors) {
            const edgeKey = [node, neighbor].sort().join("<-->"); // Standardized key

            if (adjList[neighbor]?.includes(node)) {
                // **Bidirectional → Use `<-->` (only if not processed)**
                if (!processedEdges.has(edgeKey)) {
                    undirectedEdges.push(neighbor);
                    processedEdges.add(edgeKey);
                }
            } else {
                // **One-way → Use `-->`**
                directedEdges.push(neighbor);
            }
        }

        // Group and output directed edges
        for (let i = 0; i < directedEdges.length; i += groupLimit) {
            const batch = directedEdges.slice(i, i + groupLimit).join(" & ");
            mermaid += `    ${node} --> ${batch}\n`;
        }

        // Group and output undirected edges
        for (let i = 0; i < undirectedEdges.length; i += groupLimit) {
            const batch = undirectedEdges.slice(i, i + groupLimit).join(" & ");
            mermaid += `    ${node} <--> ${batch}\n`;
        }
    }

    mermaid += `\`\`\``;
    return mermaid;
}


// Example usage
const mermaidDiagram1 = `
graph LR
    A((A))
    B((B))
    C((C))
    D((D))

    A --> B & C
    C --> D
    D --> B & A
`;

// Example usage

const mermaidDiagram2 = `
graph LR
    A((A))
    B((B))
    C((C))
    D((D))
    E((E))

    A --> B & C
    C --> D
    D --> B & A
    B --- E
    D <--> E
`;


console.log("diagram1", parseMermaidToAdjacencyList(mermaidDiagram1));
console.log("diagram2", parseMermaidToAdjacencyList(mermaidDiagram2));

const graph: AdjacencyList = {
    A: ["B", "C"],
    B: ["D"],
    C: ["D"],
    D: ["A"]
};

console.log(adjacencyListToMermaid(graph, "LR"));

const directedGraph: Record<string, string[]> = {
    A: ["B"],
    B: ["C", "D"],
    C: [],
    D: []
  };

const undirectedGraph: Record<string, string[]> = {
A: ["B", "C"],
B: ["A", "D"],
C: ["A", "D"],
D: ["B", "C"]
};

const disconnectedGraph: Record<string, string[]> = {
    A: ["B"],
    B: ["A"],
    C: ["D"],
    D: ["C"]
};

const cyclicDirectedGraph: Record<string, string[]> = {
    A: ["B"],
    B: ["C"],
    C: ["D"],
    D: ["B"] // Cycle from D → B → C → D
};

const largeGraph: Record<string, string[]> = {
    A: ["B", "C", "D"],
    B: ["E", "F"],
    C: ["G", "H"],
    D: ["I"],
    E: ["J"],
    F: ["J", "K"],
    G: ["L"],
    H: ["M"],
    I: ["N"],
    J: ["O"],
    K: [],
    L: ["P"],
    M: [],
    N: ["P"],
    O: [],
    P: []
  };

const bfsGraph = {
    A: ["B", "C"],
    B: ["A", "D", "E"],
    C: ["A", "F"],
    D: ["B"],
    E: ["B", "F"],
    F: ["C", "E", "G"],
    G: ["F"]
};

const dfsCycleGraph = {
    A: ["B", "C"],
    B: ["A", "D", "E"],
    C: ["A", "F"],
    D: ["B"],
    E: ["B", "F"],
    F: ["C", "E", "G"],
    G: ["F", "D"] // Cycle between G → D → B → E → F → G
};

const topoGraph = {
    A: ["B", "C"],
    B: ["D"],
    C: ["D"],
    D: ["E"],
    E: []
};

const dijkstraGraph = {
    A: { B: 4, C: 1 },
    B: { A: 4, D: 5, E: 8 },
    C: { A: 1, F: 6 },
    D: { B: 5, G: 3 },
    E: { B: 8, F: 2 },
    F: { C: 6, E: 2, G: 7 },
    G: { D: 3, F: 7 }
};

const bellmanFordGraph = {
    A: { B: 6, C: 7 },
    B: { D: 5, E: -4, C: 8 },
    C: { E: 9, F: 3 },
    D: { B: -2 },
    E: { D: 7, F: -2 },
    F: {}
};

const sccGraph = {
    A: ["B"],
    B: ["C"],
    C: ["A", "D"],
    D: ["E"],
    E: ["F"],
    F: ["D"]
};

console.log("directed graph:\n", adjacencyListToMermaid(directedGraph));
console.log("undirected graph:\n", adjacencyListToMermaid(undirectedGraph));
console.log("disconnected graph:\n", adjacencyListToMermaid(disconnectedGraph));
console.log("cyclic directed graph:\n", adjacencyListToMermaid(cyclicDirectedGraph));
console.log("large graph:\n", adjacencyListToMermaid(largeGraph));
console.log("breadth-first search graph:\n", adjacencyListToMermaid(bfsGraph));
console.log("depth-first cycle search graph:\n", adjacencyListToMermaid(dfsCycleGraph));
console.log("topological sort test graph:\n", adjacencyListToMermaid(topoGraph));
console.log("strongly-connected-components test graph:\n", adjacencyListToMermaid(sccGraph));

// console.log("dijkstra's algorithm test graph:\n", adjacencyListToMermaid(dijkstraGraph));
// console.log("bellman-ford algorithm test graph: \n", adjacencyListToMermaid(dijkstraGraph));
