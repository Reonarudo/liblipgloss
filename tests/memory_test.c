#include <stdio.h>
#include <stdint.h>
#include <stdbool.h>
#include <string.h>
#include "liblipgloss.h"

void test_border_memory() {
    printf("\n=== Testing Border Memory Management ===\n");
    
    // Test single border allocation/deallocation
    printf("Testing single border...\n");
    CBorder normal = NormalBorder();
    FreeBorder(normal);
    
    // Test multiple border operations
    printf("Testing multiple border operations...\n");
    for (int i = 0; i < 100; i++) {
        CBorder border;
        switch (i % 7) {
            case 0: border = NormalBorder(); break;
            case 1: border = RoundedBorder(); break;
            case 2: border = DoubleBorder(); break;
            case 3: border = ThickBorder(); break;
            case 4: border = BlockBorder(); break;
            case 5: border = InnerHalfBlockBorder(); break;
            default: border = OuterHalfBlockBorder(); break;
        }
        FreeBorder(border);
    }
    
    // Test custom border memory management
    printf("Testing custom border...\n");
    CBorder custom = CreateCustomBorder(
        "+", "-", "|", "|",  // top, bottom, left, right
        "+", "+", "+", "+",  // corners
        "+", "+", "-",       // middle parts
        "+", "+"             // middle connections
    );
    FreeBorder(custom);
}

void test_style_memory() {
    printf("\n=== Testing Style Memory Management ===\n");
    
    // Test basic style creation and deletion
    printf("Testing basic style operations...\n");
    uint64_t style = NewStyle();
    FreeStyle(style);
    
    // Test style inheritance and cleanup
    printf("Testing style inheritance...\n");
    uint64_t base_style = NewStyle();
    uint64_t child_style = CopyStyle(base_style);
    FreeStyle(child_style);
    FreeStyle(base_style);
    
    // Test style with multiple properties
    printf("Testing complex style operations...\n");
    uint64_t complex_style = NewStyle();
    complex_style = StyleBold(complex_style, 1);
    complex_style = StyleItalic(complex_style, 1);
    complex_style = StyleForeground(complex_style, "#FF0000");
    complex_style = StyleBackground(complex_style, "#00FF00");
    CBorder border = RoundedBorder();
    complex_style = StyleBorder(complex_style, border);
    FreeBorder(border);
    FreeStyle(complex_style);
}

void test_string_memory() {
    printf("\n=== Testing String Memory Management ===\n");
    
    // Test string rendering and cleanup
    printf("Testing string rendering...\n");
    uint64_t style = NewStyle();
    char* rendered = StyleRender(style, "Test String");
    FreeString(rendered);
    FreeStyle(style);
    
    // Test multiple string operations
    printf("Testing multiple string operations...\n");
    for (int i = 0; i < 100; i++) {
        uint64_t temp_style = NewStyle();
        char* str = StyleRender(temp_style, "Memory Test");
        FreeString(str);
        FreeStyle(temp_style);
    }
}

void test_color_memory() {
    printf("\n=== Testing Color Memory Management ===\n");
    
    // Test basic color operations
    printf("Testing color operations...\n");
    uint64_t style = NewStyle();
    
    // Test foreground colors
    style = StyleForeground(style, "#FF0000");
    style = StyleForeground(style, "#00FF00");
    style = StyleForeground(style, "#0000FF");
    
    // Test background colors
    style = StyleBackground(style, "#FF0000");
    style = StyleBackground(style, "#00FF00");
    style = StyleBackground(style, "#0000FF");
    
    FreeStyle(style);
}

void test_layout_memory() {
    printf("\n=== Testing Layout Memory Management ===\n");
    
    // Test horizontal join
    printf("Testing horizontal join...\n");
    char* h_joined = JoinHorizontal(0.5, "Left", "Right");
    FreeString(h_joined);
    
    // Test vertical join
    printf("Testing vertical join...\n");
    char* v_joined = JoinVertical(0.5, "Top", "Bottom");
    FreeString(v_joined);
    
    // Test placement
    printf("Testing placement...\n");
    char* placed = Place(20, 3, 0.5, 0.5, "Center");
    FreeString(placed);
}

void test_stress_memory() {
    printf("\n=== Stress Testing Memory Management ===\n");
    
    // Perform multiple operations in rapid succession
    printf("Performing stress test...\n");
    for (int i = 0; i < 1000; i++) {
        // Create style
        uint64_t style = NewStyle();
        
        // Add properties
        style = StyleBold(style, 1);
        style = StyleForeground(style, "#FF0000");
        
        // Create and add border
        CBorder border = RoundedBorder();
        style = StyleBorder(style, border);
        FreeBorder(border);
        
        // Render string
        char* rendered = StyleRender(style, "Stress Test");
        FreeString(rendered);
        
        // Cleanup style
        FreeStyle(style);
    }
}

void print_memory_status() {
    printf("\n=== Memory Status ===\n");
    char* leaks = GetMemoryLeaks();
    printf("%s\n", leaks);
    FreeString(leaks);
}

// Tables

void test_table_memory() {
    printf("\n=== Testing Table Memory Management ===\n");
    uint64_t table = NewTable();
    
    char* headers[] = {"Column1", "Column2"};
    TableAddHeaders(table, headers, 2);
    
    for (int i = 0; i < 100; i++) {
        char row[2][20];
        sprintf(row[0], "Row %d", i);
        sprintf(row[1], "Data %d", i);
        char* rowPtrs[] = {row[0], row[1]};
        TableAddRow(table, rowPtrs, 2);
    }
    
    char* result = RenderTable(table);
    printf("%s\n", result);
    FreeString(result);
    FreeTable(table);
}

// Lists

void test_list_memory() {
    printf("\n=== Testing List Memory Management ===\n");
    uint64_t list = NewList();
    
    for (int i = 0; i < 100; i++) {
        char item[20];
        sprintf(item, "Item %d", i);
        ListAddItem(list, item);
    }
    
    char* result = RenderList(list);
    printf("%s\n", result);
    FreeString(result);
    FreeList(list);
}

// Trees

void test_tree_memory() {
    printf("\n=== Testing Tree Memory Management ===\n");
    uint64_t tree = NewTree();
    
    for (int i = 0; i < 50; i++) {
        char node[20];
        sprintf(node, "Branch %d", i);
        uint64_t branch = NewTree();
        TreeAddChildValue(branch, node);
        for (int j = 0; j < 5; j++) {
            char leaf[20];
            sprintf(leaf, "Leaf %d-%d", i, j);
            TreeAddChildValue(branch, leaf);
        }
        TreeAddChildTree(tree, branch);
    }
    
    char* result = RenderTree(tree);
    printf("%s\n", result);
    FreeString(result);
    FreeTree(tree);
}

int main() {
    // Set log level to debug to track allocations
    SetLogLevel(3);  // LogLevelDebug
    
    // Run individual tests
    test_border_memory();
    test_style_memory();
    test_string_memory();
    test_color_memory();
    test_layout_memory();
    test_table_memory();
    test_list_memory();
    
    // Run stress test
    test_stress_memory();
    
    // Check final memory status
    print_memory_status();
    
    printf("\n=== Memory Tests Completed ===\n");
    return 0;
}