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

type AdjacencyList<T> = Map<T, [T, number][]>

function adjacencyListToMermaid<T>(adjList: AdjacencyList<T>, direction: "TD" | "LR" | "RL" | "BT" = "TD"): string {
    let mermaidLines: string[] = [`\`\`\`mermaid`, `graph ${direction}`];

    // Extract all unique nodes
    const nodes = new Set<T>();
    let isWeighted = false;
    for (const [node, edges] of adjList) {
        nodes.add(node);
        for (const [neighbor, weight] of edges) {
            nodes.add(neighbor);
            if (weight !== 0) isWeighted = true;
        }
    }

    // Define nodes at the top
    for (const node of nodes) {
        mermaidLines.push(`    ${node}(( ${node} ))`);
    }

    // Process edges
    const processedEdges = new Set<string>();
    for (const [node, edges] of adjList) {
        for (const [neighbor, weight] of edges) {
            let arrow: string;
            const edgeKey = `${node}-${neighbor}`;
            const reverseEdgeKey = `${neighbor}-${node}`;

            if (adjList.get(neighbor)?.some(([n, w]) => n === node && w === weight)) {
                if (!processedEdges.has(reverseEdgeKey)) {
                    arrow = `<-->`;
                    processedEdges.add(edgeKey);
                    processedEdges.add(reverseEdgeKey);
                } else {
                    continue; // Avoid duplicate bidirectional edges
                }
            } else {
                arrow = `-->`;
                processedEdges.add(edgeKey);
            }

            const weightLabel = isWeighted ? `|${weight ?? 0}| ` : "";
            mermaidLines.push(`    ${node} ${arrow} ${weightLabel}${neighbor}`);
        }
    }

    mermaidLines.push(`\`\`\``);
    return mermaidLines.join("\n");
}

// Example usage
const weightedGraph: AdjacencyList<string> = new Map([
    ["A", [["B", 4], ["C", undefined]]],
    ["B", [["D", 5], ["E", 3]]],
    ["C", [["F", undefined]]],
    ["E", [["F", 7]]],
    ["F", [["G", 6]]],
]);
console.log(adjacencyListToMermaid(weightedGraph, "LR"));
