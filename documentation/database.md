Database Schemas for the ecommerce.

Single-Inheritance style using mongo db.

Detail property can have different things.


In the following approach the index is the key to access the metadata 

```json

{
  sku  : "11ed124980e",
  category : "Ref to Category",
  description : "the description",
  brand : "Kilichips",
  name : "Kilichips Picositas Zanahoria",
  thumbnails : {
    small_devices : "https://mycdn.com/zanahoria.jpg"
  }
  shipping : {
    weights : "200m",
    dimensions : 
      {
        width: 10,
        height: 10,
        depth: 1
      },
  },

  details : {
    content: "200mg",

  }
}

```
```json
{
  "sku"  : "11ed124980e",
  "category_id" : "Ref to Category",
  "description" : "the description",
  "brand" : "Kilichips",
  "name" : "Kilichips Zanahoria",
  "thumbnails" : "http://mycdn.com/zanahoria_main.jpg",
  "variants": [
    {
      "shipping" : {
        "weight" : 200,
        "dimension" : {
            "width": 10,
            "height": 10,
            "depth": 1
          }
      },
      "thumbnails" : {
          "320_780px"   : "https://mycdn.com/zanahoria_320_780px.jpg",
          "780_1024px"  : "https://mycdn.com/zanahoria_780_1024.jpg",
          "1024_1824px" : "https://mycdn.com/zanahoria_780_1024.jpg"
      },
      "details" : {
        "flavor" :"",
        "sodium" : 300
      },
      "pricing": {
        "list": 1200,
        "retail": 1100,
        "savings": 100,
        "pct_savings": 8
      }
    }
  ]
}
```


```json
{
  sku: "00e8da9b",
  type: "Audio Album",
  title: "A Love Supreme",
  description: "by John Coltrane",
  asin: "B0000A118M",

  shipping: {
    weight: 6,
    dimensions: {
      width: 10,
      height: 10,
      depth: 1
    },
  },

  pricing: {
    list: 1200,
    retail: 1100,
    savings: 100,
    pct_savings: 8
  },

  details: {
    title: "A Love Supreme [Original Recording Reissued]",
    artist: "John Coltrane",
    genre: [ "Jazz", "General" ],
        ...
    tracks: [
      "A Love Supreme, Part I: Acknowledgement",
      "A Love Supreme, Part II: Resolution",
      "A Love Supreme, Part III: Pursuance",
      "A Love Supreme, Part IV: Psalm"
    ],
  },
}
```


Another aproach is to generate

* Product collection
```json
{
	"_id": "30671", //main item ID
	"department": "Shoes",
	"category": "Shoes/Women/Pumps",
	"brand": "Calvin Klein",
	"thumbnail": "http://cdn.../pump.jpg",
	"title": "Evening Platform Pumps",
	"description": "Perfect for a casual night out or a formal event.",
	"style": "Designer",
	…
}
```
Another Variant Data Model also is important on our product
catalog.
Variation collection.

```json
{
	"_id": "93284847362823", //variant sku
	"itemId": "30671", //references the main item
	"size": 6.0,
	"color": "red"
	…
}
```

A product has many product variants.
On Udemy video the schema is:

### Category:

```json
{
  "id": "30671",
  "name" : "zanahoria",
  "description" : "desc....ription",
  "filterOptions" : ""
}
```


### Product:
```json
{
  "category_id" : "Ref of Category",
  "name" : "zanahoria",
  "description" : "description",
  "feature" : "null",
}
```

### Variant
```json
{
  "size" : {
    "width": 10,
    "height": 10,
    "depth": 1
  },

}
```