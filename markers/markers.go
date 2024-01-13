package markers

// The JFIF file markers

const JFIF_BYTE_0 uint16 = 0x00


// JPEG-JFIF File Markers
//
// Refer to ITU-T.81 (09/92), page 32
const JFIF_BYTE_FF uint16    = 0xFF; // All markers start with this as the MSB                  
const JFIF_SOF0    uint16    = 0xC0; // Start of Frame 0, Baseline DCT                           
const JFIF_SOF1    uint16    = 0xC1; // Start of Frame 1, Extended Sequential DCT               
const JFIF_SOF2    uint16    = 0xC2; // Start of Frame 2, Progressive DCT                       
const JFIF_SOF3    uint16    = 0xC3; // Start of Frame 3, Lossless (Sequential)                 
const JFIF_DHT     uint16    = 0xC4; // Define Huffman Table                                    
const JFIF_SOF5    uint16    = 0xC5; // Start of Frame 5, Differential Sequential DCT           
const JFIF_SOF6    uint16    = 0xC6; // Start of Frame 6, Differential Progressive DCT          
const JFIF_SOF7    uint16    = 0xC7; // Start of Frame 7, Differential Loessless (Sequential)   
const JFIF_SOF9    uint16    = 0xC9; // Extended Sequential DCT, Arithmetic Coding              
const JFIF_SOF10   uint16    = 0xCA; // Progressive DCT, Arithmetic Coding                      
const JFIF_SOF11   uint16    = 0xCB; // Lossless (Sequential), Arithmetic Coding                
const JFIF_SOF13   uint16    = 0xCD; // Differential Sequential DCT, Arithmetic Coding          
const JFIF_SOF14   uint16    = 0xCE; // Differential Progressive DCT, Arithmetic Coding         
const JFIF_SOF15   uint16    = 0xCF; // Differential Lossless (Sequential), Arithmetic Coding   
const JFIF_SOI     uint16    = 0xD8; // Start of Image                                          
const JFIF_EOI     uint16    = 0xD9; // End of Image                                            
const JFIF_SOS     uint16    = 0xDA; // Start of Scan                                           
const JFIF_DQT     uint16    = 0xDB; // Define Quantization Table
const JFIF_APP0    uint16    = 0xE0; // Application Segment 0, JPEG-JFIF Image
const JFIF_COM     uint16    = 0xFE; // Comment
