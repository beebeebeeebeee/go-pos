type Item = {
  ID: number;
  Name: string;
  Price: number;
};

declare module "@wails/models" {
  type entity = {
    Item: Item;
  };
}
