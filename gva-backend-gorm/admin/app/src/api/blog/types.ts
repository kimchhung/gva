export interface Blog {
  id: number;
  createdAt: Date;
  [key: string]: any;
}

export interface CreateBlog {
  [key: string]: any;
}

export interface UpdateBlog {
  [key: string]: any;
}

export interface UpdatePartialBlog {
  [key: string]: any;
}
