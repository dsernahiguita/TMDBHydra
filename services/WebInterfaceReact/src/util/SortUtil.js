/**
* Compare util
* help call to sort multidimensional array by columns (sort numbers and string)
*/
export default class SortUtil {
  /**
  * Multi sort recursive
  */
  static multisortRecursive(a, b, columns, orderBy, index) {
    const direction = orderBy[index] === 'DESC' ? 1 : 0;


    const isNumeric = !Number.isNaN(+a[columns[index]] - +b[columns[index]]);
    let x = '';
    let y = '';

    if (a[columns[index]] !== null && a[columns[index]] !== undefined) {
      x = isNumeric ? +a[columns[index]] : a[columns[index]].toLowerCase();
    }
    if (b[columns[index]] !== null && a[columns[index]] !== undefined) {
      y = isNumeric ? +b[columns[index]] : b[columns[index]].toLowerCase();
    }

    if (x < y) {
      return direction === 0 ? -1 : 1;
    }


    if (x === y) {
      return columns.length - 1 > index
        ? SortUtil.multisortRecursive(a, b, columns, orderBy, index + 1)
        : 0;
    }

    return direction === 0 ? 1 : -1;
  }

  /**
  * Function to sort multidimensional array
  *
  * @param {array} [arr] Source array
  * @param {array} [columns] List of columns to sort
  * @param {array} [order_by] List of directions (ASC, DESC)
  * @returns {array}
  */
  static multisort(arr, columns, orderBy) {
    let columnsSort = columns;
    let orderBySort = orderBy;
    if (typeof columnsSort === 'undefined') {
      columnsSort = [];
      for (let x = 0; x < arr[0].length; x += 1) {
        columnsSort.push(x);
      }
    }

    if (typeof orderBySort === 'undefined') {
      orderBySort = [];
      for (let x = 0; x < arr[0].length; x += 1) {
        orderBySort.push('ASC');
      }
    }

    return arr.sort((a, b) => SortUtil.multisortRecursive(a, b, columnsSort, orderBySort, 0));
  }
}
