package com.adidas;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class AdidasAsciiLogo {

    private String repeat(String str, int count) {
        return count > 0 ?  repeat(str, count - 1) + str: "";
    }

    public String createAdidasAsciiLogo(int width) {
        if (width < 2) {
            throw new IllegalArgumentException("minimum width is 2");
        }

        StringBuilder sb = new StringBuilder();
        int numStripes = 3;
        int smallHeight = (int) Math.round(Math.sqrt(width));
        int totalHeight = smallHeight * numStripes;

        // Some vars to speed up things a little bit
        String stripe = this.repeat("@", width);
        String separator = this.repeat(" ", smallHeight);
        List<String> stripes = new ArrayList<String>();
        for (int i = 0; i < numStripes; i++) {
            stripes.add(stripe);
        }

        for (int h = 0; h < totalHeight; h++) {
            int offset = h % smallHeight;
            int skip = numStripes - 1 - (h / smallHeight);

            sb.append(this.repeat(" ", skip * width + offset));
            sb.append(String.join(separator, stripes.subList(skip, numStripes)));

            // Java Test expects carriage return on all lines
            sb.append("\n");
        }

        return sb.toString();
    }
}
