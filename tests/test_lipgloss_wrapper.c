#include <stdio.h>
#include <stdint.h>
#include <stdbool.h>
#include "liblipgloss.h"

void test_basic_utilities() {
    printf("\n=== Testing Basic Utilities ===\n");
    
    // Initialize renderer
    DefaultRenderer();
    
    char* profile = ColorProfile();
    printf("Color Profile: %s\n", profile);
    FreeString(profile);
    
    printf("Has Dark Background: %s\n", HasDarkBackground() ? "true" : "false");
    
    const char* test_str = "Hello\nWorld";
    printf("Height of multiline string: %d\n", Height((char*)test_str));
    printf("Width of multiline string: %d\n", Width((char*)test_str));
}

void test_borders() {
    printf("\n=== Testing Borders ===\n");
    
    // Test all border types
    printf("\n--- Border Types ---\n");
    
    // Normal border
    CBorder normal = NormalBorder();
    uint64_t normal_style = NewStyle();
    uint64_t with_normal = StyleBorder(normal_style, normal);
    char* normal_text = StyleRender(with_normal, "Normal Border");
    printf("%s\n\n", normal_text);
    FreeString(normal_text);
    FreeStyle(with_normal);
    FreeBorder(normal);

    // Rounded border
    CBorder rounded = RoundedBorder();
    uint64_t rounded_style = NewStyle();
    uint64_t with_rounded = StyleBorder(rounded_style, rounded);
    char* rounded_text = StyleRender(with_rounded, "Rounded Border");
    printf("%s\n\n", rounded_text);
    FreeString(rounded_text);
    FreeStyle(with_rounded);
    FreeBorder(rounded);

    // Double border
    CBorder double_b = DoubleBorder();
    uint64_t double_style = NewStyle();
    uint64_t with_double = StyleBorder(double_style, double_b);
    char* double_text = StyleRender(with_double, "Double Border");
    printf("%s\n\n", double_text);
    FreeString(double_text);
    FreeStyle(with_double);
    FreeBorder(double_b);

    // Thick border
    CBorder thick = ThickBorder();
    uint64_t thick_style = NewStyle();
    uint64_t with_thick = StyleBorder(thick_style, thick);
    char* thick_text = StyleRender(with_thick, "Thick Border");
    printf("%s\n\n", thick_text);
    FreeString(thick_text);
    FreeStyle(with_thick);
    FreeBorder(thick);

    // Hidden border
    CBorder hidden = HiddenBorder();
    uint64_t hidden_style = NewStyle();
    uint64_t with_hidden = StyleBorder(hidden_style, hidden);
    char* hidden_text = StyleRender(with_hidden, "Hidden Border");
    printf("%s\n\n", hidden_text);
    FreeString(hidden_text);
    FreeStyle(with_hidden);
    FreeBorder(hidden);

    // Test border with colors
    printf("\n--- Colored Borders ---\n");
    CBorder color_border = NormalBorder();
    uint64_t base_style = NewStyle();
    uint64_t colored_border = StyleBorderForeground(base_style, "#FF0000");
    colored_border = StyleBorder(colored_border, color_border);
    char* color_text = StyleRender(colored_border, "Red Border");
    printf("%s\n\n", color_text);
    FreeString(color_text);
    FreeStyle(colored_border);
    FreeBorder(color_border);

    // Test custom border
    printf("\n--- Custom Border ---\n");
    CBorder custom = CreateCustomBorder(
        "+", "+", "|", "|",     // top, bottom, left, right
        "+", "+", "+", "+",     // corners
        "+", "+", "-",          // middle parts
        "+", "+"                // middle connections
    );
    uint64_t custom_style = NewStyle();
    uint64_t with_custom = StyleBorder(custom_style, custom);
    char* custom_text = StyleRender(with_custom, "Custom Border");
    printf("%s\n\n", custom_text);
    FreeString(custom_text);
    FreeStyle(with_custom);
    FreeBorder(custom);

    // Test border sizes
    printf("\n--- Border Sizes ---\n");
    CBorder size_test = NormalBorder();
    printf("Top border size: %d\n", GetTopSize(size_test));
    printf("Bottom border size: %d\n", GetBottomSize(size_test));
    printf("Left border size: %d\n", GetLeftSize(size_test));
    printf("Right border size: %d\n", GetRightSize(size_test));
    FreeBorder(size_test);

    // Clean up base styles
    FreeStyle(normal_style);
    FreeStyle(rounded_style);
    FreeStyle(double_style);
    FreeStyle(thick_style);
    FreeStyle(hidden_style);
    FreeStyle(base_style);
    FreeStyle(custom_style);
}

void test_text_formatting() {
    printf("\n=== Testing Text Formatting ===\n");
    uint64_t base_style = NewStyle();
    
    // Test individual styles
    printf("\n--- Individual Styles ---\n");
    
    uint64_t bold_style = StyleBold(base_style, 1);
    char* bold_text = StyleRender(bold_style, "Bold Text");
    printf("Bold: %s\n", bold_text);
    FreeString(bold_text);
    FreeStyle(bold_style);

    uint64_t italic_style = StyleItalic(base_style, 1);
    char* italic_text = StyleRender(italic_style, "Italic Text");
    printf("Italic: %s\n", italic_text);
    FreeString(italic_text);
    FreeStyle(italic_style);

    uint64_t underline_style = StyleUnderline(base_style, 1);
    char* underline_text = StyleRender(underline_style, "Underlined Text");
    printf("Underline: %s\n", underline_text);
    FreeString(underline_text);
    FreeStyle(underline_style);

    uint64_t strikethrough_style = StyleStrikethrough(base_style, 1);
    char* strike_text = StyleRender(strikethrough_style, "Strikethrough Text");
    printf("Strikethrough: %s\n", strike_text);
    FreeString(strike_text);
    FreeStyle(strikethrough_style);

    uint64_t faint_style = StyleFaint(base_style, 1);
    char* faint_text = StyleRender(faint_style, "Faint Text");
    printf("Faint: %s\n", faint_text);
    FreeString(faint_text);
    FreeStyle(faint_style);

    uint64_t blink_style = StyleBlink(base_style, 1);
    char* blink_text = StyleRender(blink_style, "Blinking Text");
    printf("Blink: %s\n", blink_text);
    FreeString(blink_text);
    FreeStyle(blink_style);

    uint64_t reverse_style = StyleReverse(base_style, 1);
    char* reverse_text = StyleRender(reverse_style, "Reverse Text");
    printf("Reverse: %s\n", reverse_text);
    FreeString(reverse_text);
    FreeStyle(reverse_style);

    // Test style combinations
    printf("\n--- Style Combinations ---\n");
    
    // Bold + Italic
    uint64_t bold_italic = StyleBold(StyleItalic(base_style, 1), 1);
    char* bi_text = StyleRender(bold_italic, "Bold and Italic");
    printf("Bold + Italic: %s\n", bi_text);
    FreeString(bi_text);
    FreeStyle(bold_italic);

    // Bold + Underline + Colored
    uint64_t color_style = StyleForeground(base_style, "#FF0000");
    uint64_t bold_under_color = StyleBold(StyleUnderline(color_style, 1), 1);
    char* buc_text = StyleRender(bold_under_color, "Bold, Underlined and Red");
    printf("Bold + Underline + Red: %s\n", buc_text);
    FreeString(buc_text);
    FreeStyle(bold_under_color);
    FreeStyle(color_style);

    // Italic + Strikethrough + Colored
    uint64_t blue_style = StyleForeground(base_style, "#0000FF");
    uint64_t italic_strike_color = StyleItalic(StyleStrikethrough(blue_style, 1), 1);
    char* isc_text = StyleRender(italic_strike_color, "Italic, Strikethrough and Blue");
    printf("Italic + Strikethrough + Blue: %s\n", isc_text);
    FreeString(isc_text);
    FreeStyle(italic_strike_color);
    FreeStyle(blue_style);

    // Complex combinations
    printf("\n--- Complex Combinations ---\n");

    // Bold + Italic + Underline + Color
    uint64_t green_style = StyleForeground(base_style, "#00FF00");
    uint64_t complex1 = StyleBold(StyleItalic(StyleUnderline(green_style, 1), 1), 1);
    char* complex1_text = StyleRender(complex1, "Bold, Italic, Underline and Green");
    printf("Complex 1: %s\n", complex1_text);
    FreeString(complex1_text);
    FreeStyle(complex1);
    FreeStyle(green_style);

    // Faint + Italic + Colored Background
    uint64_t bg_style = StyleBackground(base_style, "#FFFF00");
    uint64_t complex2 = StyleFaint(StyleItalic(bg_style, 1), 1);
    char* complex2_text = StyleRender(complex2, "Faint, Italic with Yellow Background");
    printf("Complex 2: %s\n", complex2_text);
    FreeString(complex2_text);
    FreeStyle(complex2);
    FreeStyle(bg_style);

    FreeStyle(base_style);
}

void test_colors() {
    printf("\n=== Testing Colors ===\n");
    uint64_t base_style = NewStyle();
    
    // Test basic colors using hex values
    printf("\n--- Basic Colors (Hex) ---\n");
    const char* colors[] = {"#FF0000", "#00FF00", "#0000FF", "#FFFF00", "#FF00FF", "#00FFFF"};
    const char* color_names[] = {"Red", "Green", "Blue", "Yellow", "Magenta", "Cyan"};
    
    for (int i = 0; i < 6; i++) {
        uint64_t color_style = StyleForeground(base_style, (char*)colors[i]);
        char* colored_text = StyleRender(color_style, (char*)color_names[i]);
        printf("%s  ", colored_text);
        FreeString(colored_text);
        FreeStyle(color_style);
    }
    printf("\n");

    // Test background colors
    printf("\n--- Background Colors ---\n");
    const char* bg_colors[] = {"#FF0000", "#00FF00", "#0000FF"};
    
    for (int i = 0; i < 3; i++) {
        uint64_t bg_style = StyleBackground(base_style, (char*)bg_colors[i]);
        uint64_t combined = StyleForeground(bg_style, "#FFFFFF");
        char* bg_text = StyleRender(combined, "  BG  ");
        printf("%s  ", bg_text);
        FreeString(bg_text);
        FreeStyle(combined);
        FreeStyle(bg_style);
    }
    printf("\n");

    // Test color combinations
    printf("\n--- Color Combinations ---\n");
    const char* fg_colors[] = {"#FFFFFF", "#000000", "#FFFFFF"};
    const char* bg_colors_combo[] = {"#FF0000", "#FFFF00", "#0000FF"};
    const char* combo_texts[] = {"White/Red", "Black/Yellow", "White/Blue"};
    
    for (int i = 0; i < 3; i++) {
        uint64_t fg_style = StyleForeground(base_style, (char*)fg_colors[i]);
        uint64_t combo_style = StyleBackground(fg_style, (char*)bg_colors_combo[i]);
        char* combo_text = StyleRender(combo_style, (char*)combo_texts[i]);
        printf("%s  ", combo_text);
        FreeString(combo_text);
        FreeStyle(combo_style);
        FreeStyle(fg_style);
    }
    printf("\n");

    FreeStyle(base_style);
}

void test_position() {
    printf("\n=== Testing Position and Layout ===\n");
    
    char* h_joined = JoinHorizontal(PositionLeft(), "Left", "Right");
    printf("Horizontal Join: %s\n", h_joined);
    FreeString(h_joined);

    char* v_joined = JoinVertical(PositionTop(), "Top", "Bottom");
    printf("Vertical Join:\n%s\n", v_joined);
    FreeString(v_joined);

    char* placed = Place(20, 3, PositionCenter(), PositionCenter(), "Center");
    printf("Placed Text (20x3 centered):\n%s\n", placed);
    FreeString(placed);
}

// Tables
void test_table() {
    printf("\n=== Testing Table Rendering ===\n");
    uint64_t table = NewTable();
    
    char* headers[] = {"Name", "Age", "City"};
    TableAddHeaders(table, headers, 3);
    
    char* row1[] = {"Alice", "30", "New York"};
    TableAddRow(table, row1, 3);
    
    char* row2[] = {"Bob", "25", "San Francisco"};
    TableAddRow(table, row2, 3);
    
    char* result = RenderTable(table);
    printf("%s\n", result);
    FreeString(result);
    FreeTable(table);
}

void test_table_borders() {
    printf("\n=== Testing Table Borders ===\n");
    uint64_t table = NewTable();
    char* headers[] = {"Product", "Price"};
    TableAddHeaders(table, headers, 2);
    char* row1[] = {"Laptop", "$1000"};
    TableAddRow(table, row1, 2);
    char* row2[] = {"Phone", "$500"};
    TableAddRow(table, row2, 2);
    
    printf("Normal Border:\n");
    TableSetBorder(table, 0);
    char* normal = RenderTable(table);
    printf("%s\n", normal);
    FreeString(normal);
    
    printf("Rounded Border:\n");
    TableSetBorder(table, 1);
    char* rounded = RenderTable(table);
    printf("%s\n", rounded);
    FreeString(rounded);
    
    printf("Thick Border:\n");
    TableSetBorder(table, 2);
    char* thick = RenderTable(table);
    printf("%s\n", thick);
    FreeString(thick);
    
    FreeTable(table);
}

// Lists

void test_list() {
    printf("\n=== Testing List Rendering ===\n");
    uint64_t list = NewList();
    
    ListAddItem(list, "Apples");
    ListAddItem(list, "Bananas");
    ListAddItem(list, "Oranges");
    
    char* result = RenderList(list);
    printf("%s\n", result);
    FreeString(result);
    FreeList(list);
}

void test_list_enumerators() {
    printf("\n=== Testing List Enumerators ===\n");
    uint64_t list = NewList();
    ListAddItem(list, "Foo");
    ListAddItem(list, "Bar");
    ListAddItem(list, "Baz");
    
    printf("Bullet Enumerator:\n");
    ListSetEnumerator(list, 0);
    char* bullet = RenderList(list);
    printf("%s\n", bullet);
    FreeString(bullet);
    
    printf("Dash Enumerator:\n");
    ListSetEnumerator(list, 1);
    char* dash = RenderList(list);
    printf("%s\n", dash);
    FreeString(dash);
    
    printf("Roman Enumerator:\n");
    ListSetEnumerator(list, 4);
    char* roman = RenderList(list);
    printf("%s\n", roman);
    FreeString(roman);
    
    FreeList(list);
}

// Trees

void test_tree() {
    printf("\n=== Testing Tree Rendering ===\n");
    uint64_t tree = NewTree();
    TreeSetItemStyle(tree, "#FF0000");
    TreeAddChildValue(tree, "Foo");
    uint64_t bar = NewTree();
    TreeSetItemStyle(bar, "#00FF00");
    TreeAddChildValue(bar, "Bar");
    uint64_t qux = NewTree();
    TreeSetItemStyle(qux, "#0000FF");
    TreeAddChildValue(qux, "Qux");
    uint64_t quux = NewTree();
    TreeSetItemStyle(quux, "#FF00FF");
    TreeAddChildValue(quux, "Quux");
    TreeAddChildValue(quux, "Foo");
    TreeAddChildValue(quux, "Bar");
    TreeAddChildTree(qux, quux);
    TreeAddChildValue(qux, "Quuux");
    TreeAddChildTree(bar, qux);
    TreeAddChildTree(tree, bar);
    TreeAddChildValue(tree, "Baz");

    char* result = RenderTree(tree);
    printf("%s\n", result);
    FreeString(result);
    FreeTree(tree);
}

void test_tree_enumerators() {
    printf("\n=== Testing Tree Enumerators ===\n");
    uint64_t tree = NewTree();

    TreeAddChildValue(tree, "Foo");
    uint64_t bar = NewTree();
    TreeAddChildValue(bar, "Bar");
    uint64_t qux = NewTree();
    TreeAddChildValue(qux, "Qux");
    uint64_t quux = NewTree();
    TreeAddChildValue(quux, "Quux");
    TreeAddChildValue(quux, "Foo");
    TreeAddChildValue(quux, "Bar");
    TreeAddChildTree(qux, quux);
    TreeAddChildValue(qux, "Quuux");
    TreeAddChildTree(bar, qux);
    TreeAddChildTree(tree, bar);
    TreeAddChildValue(tree, "Baz");

    printf("Default Enumerator:\n");
    TreeSetEnumerator(tree, 0);
    char* default_enum = RenderTree(tree);
    printf("%s\n", default_enum);
    FreeString(default_enum);

    printf("Rounded Enumerator:\n");
    TreeSetEnumerator(tree, 1);
    char* rounded_enum = RenderTree(tree);
    printf("%s\n", rounded_enum);
    FreeString(rounded_enum);

    FreeTree(tree);
}

int main() {
    test_basic_utilities();
    test_text_formatting();
    test_colors();
    test_position();
    test_borders();
    test_table();
    test_table_borders();
    test_list();
    test_list_enumerators();
    test_tree();
    test_tree_enumerators();
    
    printf("\n=== All tests completed ===\n");
    return 0;
}