/**
 * Returns the adidas three stripes logo using `@` characters.
 * @param {number} width - Width of a stripe.
 * @returns {string} adidas logo.
 */
module.exports = function(width) {
  if (width < 2) {
    throw 'Error, minimum width is 2';
  }

  let str = '';
  let numStripes = 3;
  let smallHeight = Math.round(Math.sqrt(width));
  let totalHeight = smallHeight * numStripes

  // Some vars to speed up things a little bit
  let stripe = '@'.repeat(width);
  let separator = ' '.repeat(smallHeight);
  let stripes = Array(numStripes).fill(stripe);

  for (let h = 0; h < totalHeight; h++) {
    let offset = h % smallHeight;
    let skip = numStripes - 1 - Math.floor(h / smallHeight);

    str += ' '.repeat( (skip * width) + offset);
    str += stripes.slice(skip, numStripes).join(separator);

    // Tests expect no last carriage return
    if (h != totalHeight - 1) {
      str += '\n';
    }
  }

  return str;
};
