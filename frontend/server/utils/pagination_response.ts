export default interface PaginationResponse<T> {
  items: T[];
  page: number;
  size: number;
  total: number;
}
