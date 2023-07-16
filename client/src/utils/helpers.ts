import { RecentlyChattedFriends } from "../types/Chat";

export function sortByProperty(arr: RecentlyChattedFriends[], property: keyof RecentlyChattedFriends, descending: boolean): RecentlyChattedFriends[] {
  return arr.sort((a, b) => {
    const valueA = a[property];
    const valueB = b[property];

    if (valueA < valueB) {
      return descending ? 1 : -1;
    }
    if (valueA > valueB) {
      return descending ? -1 : 1;
    }
    return 0;
  });
}
